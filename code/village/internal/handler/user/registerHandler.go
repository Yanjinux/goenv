/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 01:01:52
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-25 01:31:37
 * @FilePath: \code\village\internal\handler\user\registerHandler.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

import (
	"net/http"

	"village/village/internal/logic/user"
	"village/village/internal/svc"
	"village/village/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
