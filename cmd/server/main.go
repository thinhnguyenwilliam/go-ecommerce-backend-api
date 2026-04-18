// go-ecommerce-backend-api/cmd/server/main.go
// lsof -i :8085,  kill -9 12345, kill -9 $(lsof -t -i:8085)
package main

import "github.com/thinhnguyenwilliam/go-ecommerce-backend-api/internal/initialize"

func main() {
	initialize.Run()
}
