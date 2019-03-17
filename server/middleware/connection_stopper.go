package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
)

var (
	once     sync.Once
	instance *ConnextionStopper
)

type ConnextionStopper struct {
	Semaphore chan bool
	Max       int64
}

func GetInstance() *ConnextionStopper {
	once.Do(func() {
		instance = &ConnextionStopper{}
	})
	return instance
}

func (cs *ConnextionStopper) Setup(max int64) {
	cs.Max = max
	cs.Semaphore = make(chan bool, max)
}

func (cs *ConnextionStopper) GetCurrent() int {
	return len(cs.Semaphore)
}

func (cs *ConnextionStopper) Limit() gin.HandlerFunc {
	if cs.Max <= 0 {
		logrus.Warnf("Max current connections was less than 1: %d. Set to 1 by default.", cs.Max)
		cs.Max = 1
	}
	return func(ctx *gin.Context) {
		if strings.HasPrefix("/stats", ctx.Request.RequestURI) {
			ctx.Next()
			return
		}
		var called, fulled bool
		defer func() {
			if !called && !fulled {
				<-cs.Semaphore
			}
			if r := recover(); r != nil {
				panic(r)
			}
		}()
		select {
		case cs.Semaphore <- true:
			logrus.Warnf("Adding connection: %s", ctx.Request.RequestURI)
			ctx.Next()
			called = true
			logrus.Warnf("Closing connection: %s", ctx.Request.RequestURI)
			<-cs.Semaphore
		default:
			fulled = true
			logrus.Warnf("Reached maximum amount of current connections: %d. Request rejected: %s", cs.Max, ctx.Request.RequestURI)
			ctx.Status(http.StatusBadGateway)
			ctx.Abort()
		}
	}
}
