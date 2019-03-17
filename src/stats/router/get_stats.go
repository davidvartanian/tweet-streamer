package router

import (
	"github.com/gin-gonic/gin"
	"main/src/stats/handler"
)

func SetupRoute(engine *gin.Engine) {
	engine.GET("/stats", handler.HandleStats)
}
