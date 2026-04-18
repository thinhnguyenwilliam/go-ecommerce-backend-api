// go-ecommerce-backend-api/internal/initialize/logger.go
package initialize

import (
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/global"
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(&global.Config.Logger)
	global.Logger.Info("server started")
}
