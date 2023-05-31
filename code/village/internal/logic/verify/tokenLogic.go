/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 00:59:47
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-29 00:40:26
 * @FilePath: \code\village\internal\logic\verify\tokenLogic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package verify

import (
	"context"

	"village/village/common/ctxData"
	"village/village/common/xerr"
	"village/village/internal/svc"
	"village/village/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var ValidateTokenError = xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR)

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenLogic {
	return &TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token(req *types.VerifyTokenReq) (resp *types.VerifyTokenResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxData.GetUidFromCtx(l.ctx)

	if userID == 0 {
		return nil, errors.Wrapf(ValidateTokenError, "JwtAuthLogic isPass  ParseToken err : %v", err)
	}

	return
}
