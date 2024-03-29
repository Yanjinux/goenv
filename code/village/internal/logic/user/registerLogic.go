/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 01:01:52
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-06-10 00:14:10
 * @FilePath: \goenv\code\village\internal\logic\user\registerLogic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"village/village/internal/svc"
	"village/village/internal/types"

	"village/village/common/ctxData"
	"village/village/common/tool"
	"village/village/common/xerr"

	"village/village/model"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUserAlreadyRegisterError = errors.New("该用户已被注册")

func GetUidFromCtx(ctx context.Context) int64 {
	uid, err := strconv.ParseInt(string(ctx.Value(ctxData.CtxKeyJwtUserId).(json.Number)), 10, 64)
	fmt.Printf("Get UID %v %t", err, ctx.Value(ctxData.CtxKeyJwtUserId))
	return uid
}

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *RegisterLogic) checkVerifyToken(mobile, token string) bool {
	rt, err := l.svcCtx.RedisClient.Get(fmt.Sprintf(ctxData.MobileRegistTokenKey, mobile))
	if err != nil {
		fmt.Println(rt, "|", token)

		l.Logger.Errorf("get mobile token fail %v", err)
		return false
	}
	fmt.Println(rt, "|", token)
	return rt == token
}
func (l *RegisterLogic) Register(in *types.RegisterReq) (resp *types.RegisterResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
	}
	if user != nil {
		return nil, ErrUserAlreadyRegisterError
	}

	// 校验验证码
	if !l.checkVerifyToken(in.Mobile, in.Token) {
		return nil, xerr.NewErrMsg("校验失败")
	}
	user = new(model.User)
	user.Mobile = in.Mobile
	if len(user.Nickname) == 0 {
		user.Nickname = tool.Krand(tool.KC_RAND_KIND_ALL, 8)
	}
	if len(in.Password) > 0 {
		user.Password = tool.Md5ByString(in.Password)
	}
	insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err:%v,user:%+v", err, user)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insertResult.LastInsertId err:%v,user:%+v", err, user)
	}

	if err != nil {
		return nil, err
	}

	//2、生成token.
	resp, err = GenerateToken(l.svcCtx, userId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "IdentityRpc.GenerateToken userId : %d , err:%+v", userId, err)
	}

	return resp, nil
}

// GenerateToken 创建token
func GenerateToken(ctx *svc.ServiceContext, useID int64) (*types.RegisterResp, error) {

	now := time.Now().Unix()
	accessExpire := ctx.Config.Auth.AccessExpire
	accessToken, err := getJwtToken(ctx.Config.Auth.AccessSecret, now, accessExpire, useID)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", useID, err)
	}

	return &types.RegisterResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

// getJwtToken 生成JWT token
func getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxData.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
