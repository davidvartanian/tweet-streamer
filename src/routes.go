package src

import "github.com/gin-gonic/gin"

type Router interface {
	SetupRoute(engine *gin.Engine)
}

func SetupAllRoutes(engine *gin.Engine, routes ...Router) {
	for _, r := range routes {
		r.SetupRoute(engine)
	}
}
