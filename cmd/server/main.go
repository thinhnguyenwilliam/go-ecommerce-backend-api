// go-ecommerce-backend-api/cmd/server/main.go
package main

import "github.com/thinhnguyenwilliam/go-ecommerce-backend-api/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run(":8085")
}
