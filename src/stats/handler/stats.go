package handler

import (
	"github.com/gin-gonic/gin"
	"main/server/middleware"
)

// Handle tweets sample requests
// @Summary Get statistics about the usage of this microservice
// @Router /stats [get]
// @Produce  json
// @Success 200 {object} handler.Stats
func HandleStats(ctx *gin.Context) {
	cs := middleware.GetInstance()
	stats := Stats{
		MaxConnections:     cs.Max,
		CurrentConnections: cs.GetCurrent(),
	}
	ctx.JSON(200, stats)
}
