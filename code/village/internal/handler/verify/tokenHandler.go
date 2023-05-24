package verify

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"village/village/internal/logic/verify"
	"village/village/internal/svc"
	"village/village/internal/types"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := verify.NewTokenLogic(r.Context(), svcCtx)
		resp, err := l.Token(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
