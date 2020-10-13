package handler

import (
	"net/http"

	"github.com/just-coding-0/learn_example/micro_service/zero/internal/logic"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/svc"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func echoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EchoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEchoLogic(r.Context(), ctx)
		resp, err := l.Echo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
