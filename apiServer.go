package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/api"
)

func main() {
	// start server
	router := gin.Default()
	api.SetupRouter(router)
	router.Run(":10086")
}