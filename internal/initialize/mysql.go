// go-ecommerce-backend-api/internal/initialize/mysql.go
package initialize

import (
	"fmt"
	"log"
	"time"

	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	cfg := global.Config.DBMySQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect DB:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	global.DB = db

	log.Println("MySQL connected 🚀")
}
