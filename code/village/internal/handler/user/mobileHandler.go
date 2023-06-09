package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"village/village/internal/logic/user"
	"village/village/internal/svc"
	"village/village/internal/types"
)

func MobileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserSmsVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewMobileLogic(r.Context(), svcCtx)
		resp, err := l.Mobile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
