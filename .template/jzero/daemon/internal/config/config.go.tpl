package config

import (
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

var C Config

type Config struct {
	zrpc.RpcServerConf
	Gateway gateway.GatewayConf

	{{ .APP | FirstUpper }} {{ .APP | FirstUpper }}Config
}

type {{ .APP | FirstUpper }}Config struct {
	ListenOnUnixSocket string `json:",optional"`
	GrpcMaxConns       int    `json:",default=10000"`
}
