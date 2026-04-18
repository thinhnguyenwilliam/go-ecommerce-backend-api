// go-ecommerce-backend-api/global/global.go
package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/pkg/logger"
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	DB     *gorm.DB
	Logger *logger.LoggerZap
	Rdb    *redis.Client
)
