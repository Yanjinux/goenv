package user

import (
	"context"
	"fmt"

	"village/village/common/ctxData"
	"village/village/common/tool"
	"village/village/common/xerr"
	"village/village/internal/svc"
	"village/village/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileLogic {
	return &MobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 发送注册token 当前发送当前todo
func (l *MobileLogic) Mobile(req *types.UserSmsVerifyReq) (resp *types.UserSmsVerifyResp, err error) {
	// 1.判断是否已经发送 -防止频繁Mobile
	key := fmt.Sprintf(ctxData.MobileRegistTokenKey, req.Mobile)
	exist, err := l.svcCtx.RedisClient.Exists(key)
	if err != nil {
		l.Logger.Errorf("redis exists fail %v", err)
		return nil, xerr.NewErrMsg("内部错误,请稍后再试")
	}

	if exist {
		return &types.UserSmsVerifyResp{
			Code: xerr.MSG_SEND_FREQUENCE,
			Desc: "短信发送频繁,请稍后再试",
		}, nil
	}
	// 2. 发送短信

	// 3. 记录redis
	_, err = l.svcCtx.RedisClient.SetnxEx(key, tool.RandStringBytes(6), 120)
	if err != nil {
		l.Logger.Errorf("redis set key fail %v", err)
		return nil, xerr.NewErrMsg("内部错误,请稍后再试")
	}

	return &types.UserSmsVerifyResp{
		Code: xerr.OK,
	}, nil
}
