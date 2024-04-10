package hello

import (
	"net/http"

	"github.com/jaronnie/jzero/jzerod/internal/logic/hello"
	"github.com/jaronnie/jzero/jzerod/internal/svc"
	"github.com/jaronnie/jzero/jzerod/internal/types"
	"github.com/jaronnie/jzero/jzerod/pkg/response"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloPathHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := hello.NewHelloPathLogic(r.Context(), svcCtx)
		resp, err := l.HelloPath(&req)
		response.Response(w, resp, err)
	}
}