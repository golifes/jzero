package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jzero-io/jzero/app/internal/config"
	"github.com/jzero-io/jzero/app/internal/handler"
	"github.com/jzero-io/jzero/app/internal/svc"
	"github.com/jzero-io/jzero/app/middlewares"
)

func Start(cfgFile string) {
	var c config.Config
	conf.MustLoad(cfgFile, &c, conf.UseEnv())
	// set up logger
	if err := logx.SetUp(c.Log.LogConf); err != nil {
		logx.Must(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.Username,
		c.Mysql.Password,
		c.Mysql.Address,
		c.Mysql.Database)

	conn := sqlx.NewSqlConn("mysql", dsn)
	ctx := svc.NewServiceContext(c, conn)
	start(ctx)
}

func start(ctx *svc.ServiceContext) {
	// print log to console if Log.Mode is file or volume
	middlewares.PrintLogToConsole(ctx.Config)

	s := getZrpcServer(ctx.Config, ctx)

	// verify sql conn
	_, err := ctx.SqlConn.Exec("select 1 = 1")
	if err != nil {
		logx.Error(err)
	}

	middlewares.RateLimit = syncx.NewLimit(ctx.Config.Zrpc.MaxConns)
	s.AddUnaryInterceptors(middlewares.GrpcRateLimitInterceptors)

	gw := gateway.MustNewServer(ctx.Config.Gateway.GatewayConf)

	gw.Use(middlewares.WrapResponse)
	httpx.SetErrorHandler(middlewares.ErrorHandler)

	// gw add routes
	handler.RegisterMyHandlers(gw.Server, ctx)

	// gw add api routes
	handler.RegisterHandlers(gw.Server, ctx)

	// listen unix
	var unixListener net.Listener
	if ctx.Config.Gateway.ListenOnUnixSocket != "" {
		sock := ctx.Config.Gateway.ListenOnUnixSocket
		_ = os.Remove(ctx.Config.Gateway.ListenOnUnixSocket)
		unixListener, err = net.Listen("unix", sock)
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Printf("Starting unix server at %s...\n", ctx.Config.Gateway.ListenOnUnixSocket)
			if err := http.Serve(unixListener, gw); err != nil {
				return
			}
		}()
	}

	group := service.NewServiceGroup()
	group.Add(s)
	group.Add(gw)

	go func() {
		fmt.Printf("Starting rpc server at %s...\n", ctx.Config.Zrpc.ListenOn)
		fmt.Printf("Starting gateway server at %s:%d...\n", ctx.Config.Gateway.Host, ctx.Config.Gateway.Port)
		group.Start()
	}()

	signalHandler(ctx, group, unixListener)
}

func signalHandler(ctx *svc.ServiceContext, serviceGroup *service.ServiceGroup, unixListener net.Listener) {
	// signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("Waiting 1 second...\nStopping rpc server and gateway server")
			time.Sleep(time.Second)
			serviceGroup.Stop()
			if ctx.Config.Gateway.ListenOnUnixSocket != "" {
				fmt.Println("Stopping unix server")
				unixListener.Close()
				_ = os.Remove(ctx.Config.Gateway.ListenOnUnixSocket)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
