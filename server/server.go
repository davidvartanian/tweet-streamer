package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"main/server/middleware"
	"main/src"
	"main/src/stats"
	"main/src/streamer"
	"os"
	"strconv"
)

// @title Twitter Realtime Streamer
// @version 0.1
// @description This microservice creates Server Sent Event streams given a query string
func GetServer() *gin.Engine {
	if os.Getenv("DEBUG") != "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	engine.Use(logger.SetLogger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())
	max, err := strconv.ParseInt(os.Getenv("MAX_CONCURRENCY"), 0, 64)
	if err != nil {
		logrus.Warnf("Error trying to parse integer from string: %v. Setting max=1 by default", err)
		max = 1
	}
	cs := middleware.GetConnectionStopper()
	cs.Setup(max)
	engine.Use(cs.Limit())
	src.SetupAllRoutes(engine, new(streamer.Handler), new(stats.Handler))
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return engine
}
