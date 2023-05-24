/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 01:01:52
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-25 01:47:43
 * @FilePath: \goenv\code\village\internal\logic\user\registerLogic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

import (
	"context"
	"fmt"
	"time"

	"village/village/internal/svc"
	"village/village/internal/types"

	"village/village/common/tool"
	"village/village/common/xerr"
	"village/village/model"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const CacheUserTokenKey = "user_token:%d"

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUserAlreadyRegisterError = errors.New("该用户已被注册")

// 从ctx获取uid
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	uid, _ := ctx.Value(CtxKeyJwtUserId).(int64)
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

func (l *RegisterLogic) Register(in *types.RegisterReq) (resp *types.RegisterResp, err error) {
	fmt.Println("1")
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		fmt.Println(err)

		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
	}
	if user != nil {
		fmt.Println("3")

		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "用户已经存在 mobile:%s,err:%v", in.Mobile, err)
	}
	fmt.Println("2")

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
	fmt.Println("3")

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insertResult.LastInsertId err:%v,user:%+v", err, user)
	}

	if err != nil {
		fmt.Println("5")

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
	accessExpire := ctx.Config.JwtAuth.AccessExpire
	accessToken, err := getJwtToken(ctx.Config.JwtAuth.AccessSecret, now, accessExpire, useID)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", useID, err)
	}

	// 存入redis.
	userTokenKey := fmt.Sprintf(CacheUserTokenKey, useID)
	err = ctx.RedisClient.Setex(userTokenKey, accessToken, int(accessExpire))
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "SetnxEx err userId:%d, err:%v", useID, err)
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
	claims[CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
