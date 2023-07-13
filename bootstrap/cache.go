// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"github.com/lihc520/gohub/pkg/cache"
	"github.com/lihc520/gohub/pkg/config"
)

// SetupCache 缓存
func SetupCache() {

	// 初始化缓存专用的 redis Client,使用专属缓存 DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("reids.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}
