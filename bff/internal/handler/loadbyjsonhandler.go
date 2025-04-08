package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"start/bff/internal/logic"
	"start/bff/internal/svc"
	"start/bff/internal/types"
)

func LoadByjsonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoadByjsonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoadByjsonLogic(r.Context(), svcCtx)
		resp, err := l.LoadByjson(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
