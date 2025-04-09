package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"naming/bff/internal/logic"
	"naming/bff/internal/svc"
	"naming/bff/internal/types"
)

func LoadByjsonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoadByjsonReq
		bs, e := ioutil.ReadAll(r.Body)
		if e != nil {
			httpx.ErrorCtx(r.Context(), w, e)
			return
		}
		req.LoadJson = string(bs)
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//	return
		//}

		l := logic.NewLoadByjsonLogic(r.Context(), svcCtx)
		resp, err := l.LoadByjson(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
