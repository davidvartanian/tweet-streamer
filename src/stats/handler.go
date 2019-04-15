package stats

import (
	"github.com/gin-gonic/gin"
	"main/server/middleware"
)

type Handler struct{}

func (h *Handler) SetupRoute(engine *gin.Engine) {
	engine.GET("/stats", h.handleStats)
}

// Handle tweets sample requests
// @Summary Get statistics about the usage of this microservice
// @Router /stats [get]
// @Produce  json
// @Success 200 {object} handler.Stats
func (h *Handler) handleStats(ctx *gin.Context) {
	cs := middleware.GetConnectionStopper()
	stats := Stats{
		MaxConnections:     cs.Max,
		CurrentConnections: cs.GetCurrent(),
	}
	ctx.JSON(200, stats)
}
