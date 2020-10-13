package handler

import (
	"net/http"

	"github.com/just-coding-0/learn_example/micro_service/zero/internal/logic"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/svc"
	"github.com/just-coding-0/learn_example/micro_service/zero/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getEchoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetEchoStatsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetEchoLogic(r.Context(), ctx)
		resp, err := l.GetEcho(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
