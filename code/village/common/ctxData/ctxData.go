package ctxData

import (
	"context"
)

// 从ctx获取uid
var CtxKeyJwtUserId = "jwtUserId"

const CacheUserTokenKey = "user_token:%d"

func GetUidFromCtx(ctx context.Context) int64 {
	uid, _ := ctx.Value(CtxKeyJwtUserId).(int64)
	return uid
}
