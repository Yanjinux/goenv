/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 00:59:47
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-25 02:09:08
 * @FilePath: \code\village\internal\svc\serviceContext.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package svc

import (
	"village/village/internal/config"
	"village/village/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	rdi, _ := redis.NewRedis(c.Redis)
	return &ServiceContext{
		Config:      c,
		UserModel:   model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		RedisClient: rdi,
	}
}
