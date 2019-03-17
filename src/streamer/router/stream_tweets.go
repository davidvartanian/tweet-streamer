package router

import (
	"github.com/gin-gonic/gin"
	"main/src/streamer/handler"
)

func SetupRoute(engine *gin.Engine) {
	group := engine.Group("/tweets")
	{
		group.GET("/stream", handler.HandleTweetStream)
		group.GET("/sample", handler.HandleTweetSample)
	}
}
