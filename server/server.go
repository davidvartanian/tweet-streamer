package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"main/server/middleware"
	statsrouter "main/src/stats/router"
	streamrouter "main/src/streamer/router"
	"os"
	"strconv"
)

// @title Twitter Realtime Streamer
// @version 0.1
// @description This microservice creates Server Sent Event streams given a query string
func GetServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(logger.SetLogger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())
	max, err := strconv.ParseInt(os.Getenv("MAX_CONCURRENCY"), 0, 64)
	if err != nil {
		logrus.Warnf("Error trying to parse integer from string: %v. Setting max=1 by default", err)
		max = 1
	}
	cs := middleware.GetInstance()
	cs.Setup(max)
	engine.Use(cs.Limit())
	streamrouter.SetupRoute(engine)
	statsrouter.SetupRoute(engine)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return engine
}
