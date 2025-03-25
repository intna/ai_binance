package main

import (
	"ai_binance/src/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// initialize Gin
	r := gin.Default()

	// register routes
	routes.RegisterRoutes(r)

	// start server
	port := ":8080"
	if err := r.Run(port); err != nil {
		log.Fatalf("run err")
	}
}
