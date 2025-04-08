package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"start/bff/internal/logic"
	"start/bff/internal/svc"
	"start/bff/internal/types"
)

func UpdateHostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateHostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateHostLogic(r.Context(), svcCtx)
		resp, err := l.UpdateHost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
