/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 01:01:52
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-25 02:13:00
 * @FilePath: \code\village\internal\logic\user\detailLogic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

import (
	"context"

	"village/village/internal/svc"
	"village/village/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := GetUidFromCtx(l.ctx)

	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		UserInfo: types.User{
			Nickname: userInfo.Nickname,
			Mobile:   userInfo.Mobile,
			Sex:      userInfo.Sex,
			Info:     userInfo.Info,
		},
	}, nil
}
