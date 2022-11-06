package main

import (
	"10xers/configs"
	routes "10xers/routes/public"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// database connect
	configs.Connect()

	// routes
	routes.TokenRoute(router)

	router.Run()
}
