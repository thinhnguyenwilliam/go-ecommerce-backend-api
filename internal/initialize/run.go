// go-ecommerce-backend-api/internal/initialize/run.go
package initialize

import (
	"context"

	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	InitConfig()
	InitLogger()
	InitMySQL()

	// Redis
	InitRedis()
	err := global.Rdb.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, _ := global.Rdb.Get(context.Background(), "key").Result()
	global.Logger.Info("redis value", zap.String("val", val))

	// app
	r := InitRouter()
	r.Run(":8085")
}
