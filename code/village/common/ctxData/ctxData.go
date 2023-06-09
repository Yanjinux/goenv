/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-29 00:33:28
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-06-09 23:30:21
 * @FilePath: \code\village\common\ctxData\ctxData.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package ctxData

import (
	"context"
)

// 从ctx获取uid
var CtxKeyJwtUserId = "jwtUserId"

const (
	CacheUserTokenKey    = "user_token:%d"
	MobileRegistTokenKey = "register_user:%s"
)

func GetUidFromCtx(ctx context.Context) int64 {
	uid, _ := ctx.Value(CtxKeyJwtUserId).(int64)
	return uid
}
