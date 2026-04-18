// go-ecommerce-backend-api/internal/initialize/redis.go
package initialize

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/global"
)

var ctx = context.Background()

func InitRedis() {
	cfg := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Host + ":" + cfg.Port,
		Password:     cfg.Password, // nếu có
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
	})

	// test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Cannot connect Redis:", err)
		global.Logger.Error("Cannot connect Redis")
	}

	global.Rdb = rdb

	log.Println("Redis connected 🚀")
}
