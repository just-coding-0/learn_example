package handler

import (
	"net/http"

	"github.com/just-coding-0/learn_example/micro_service/zero/internal/logic"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func lastEchoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLastEchoLogic(r.Context(), ctx)
		resp, err := l.LastEcho()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
