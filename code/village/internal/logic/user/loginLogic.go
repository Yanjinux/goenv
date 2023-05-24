/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 01:01:52
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-25 02:04:39
 * @FilePath: \code\village\internal\logic\user\loginLogic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

import (
	"context"

	"village/village/common/tool"
	"village/village/common/xerr"
	"village/village/internal/svc"
	"village/village/internal/types"
	"village/village/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")
var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var userId int64
	//switch in.AuthType {
	userId, err = l.loginByMobile(req.Mobile, req.Password)
	if err != nil {
		return nil, err
	}

	//2、生成token
	tokenResp, err := GenerateToken(l.svcCtx, userId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "IdentityRpc.GenerateToken userId : %d", userId)
	}

	return &types.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil

}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	return user.Id, nil
}
