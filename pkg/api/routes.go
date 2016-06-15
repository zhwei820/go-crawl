package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	serverApi := router.Group("/api")
	{
		//serverApi.POST("/uid", GetWhichUid)

		serverApi.GET("/:type/profile/:userId", GetProfile)
		serverApi.GET("/:type/posts/:userId", GetPosts)
		serverApi.POST("/:type/uid", GetUid)
		//apiFB := serverApi.Group("/fb")
		//{
		//	apiFB.POST("/uid", GetFBUid)
		//	apiFB.GET("/profile/:userId", GetFBProfile)
		//	apiFB.GET("/posts/:userId", GetFBPosts)
		//}
		//
		//apiIG := serverApi.Group("/ig")
		//{
		//	apiIG.POST("/uid", GetIGUid)
		//	apiIG.GET("/profile/:userId", GetIGProfile)
		//	apiIG.GET("/posts/:userId", GetIGPosts)
		//}
		//
		//apiWX := serverApi.Group("/wx")
		//{
		//	apiWX.POST("/uid", GetWXUid)
		//	apiWX.GET("/profile/:userId", GetWXProfile)
		//	apiWX.GET("/posts/:userId", GetWXPosts)
		//}
		//
		//apiWB := serverApi.Group("/wb")
		//{
		//	apiWB.POST("/uid", GetWBUid)
		//	apiWB.GET("/profile/:userId", GetWBProfile)
		//	apiWB.GET("/posts/:userId", GetWBPosts)
		//}

		//serverApi.POST("/token", UpdateToken)

	}
}
