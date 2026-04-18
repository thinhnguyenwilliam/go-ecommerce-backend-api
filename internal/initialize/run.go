// go-ecommerce-backend-api/internal/initialize/run.go
package initialize

func Run() {
	InitConfig()
	InitLogger()
	InitMySQL()

	r := InitRouter()
	r.Run(":8085")
}
