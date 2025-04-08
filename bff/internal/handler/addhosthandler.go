package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"start/bff/internal/logic"
	"start/bff/internal/svc"
	"start/bff/internal/types"
)

func AddHostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddHostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddHostLogic(r.Context(), svcCtx)
		resp, err := l.AddHost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
