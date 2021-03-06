package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/llitfkitfk/cirkol/pkg/controllers"
)

func setupWBRouters(router *gin.RouterGroup) {
	router.GET("/api/wb/profile/:userId", controllers.GetWBProfile)
	router.GET("/api/wb/posts/:userId", controllers.GetWBPosts)

	router.POST("/api/wb/post", controllers.GetWBPostInfo)
}
