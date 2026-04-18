// go-ecommerce-backend-api/internal/initialize/config.go
package initialize

import (
	"log"

	"github.com/spf13/viper"
	"github.com/thinhnguyenwilliam/go-ecommerce-backend-api/global"
)

func InitConfig() {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatal(err)
	}
}
