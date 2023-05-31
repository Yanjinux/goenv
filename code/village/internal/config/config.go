/*
 * @Author: Yanjinux 471573617@qq.com
 * @Date: 2023-05-18 00:59:47
 * @LastEditors: Yanjinux 471573617@qq.com
 * @LastEditTime: 2023-05-29 00:57:16
 * @FilePath: \goenv\code\village\internal\config\config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Redis redis.RedisConf `json:",optional"`

	DB struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
