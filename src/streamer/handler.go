package streamer

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"main/src/streamer/service"
)

type Handler struct{}

func (h *Handler) SetupRoute(engine *gin.Engine) {
	group := engine.Group("/tweets")
	{
		group.GET("/stream", h.handleTweetStream)
		group.GET("/sample", h.handleTweetSample)
	}
}

// Handle tweets stream request
// @Summary Get a stream of Server Sent Events containing tweets given a search query string
// @Description This endpoint produces Server-Sent-Events from tweets given a search query string
// @Description Events occur in real-time according to Twitter activity
// @Router /tweets/stream [get]
// @Param q query string false "name search by q"
// @Produce json-stream
// @Success 200 "There is no response body for this endpoint. Event stream data: {\"tweet\": \"string\"}"
// @Failure 502 "Reached the maximum amount of simultaneous connections"
// @Tags tweets
func (h *Handler) handleTweetStream(ctx *gin.Context) {
	var query Query
	err := ctx.BindQuery(&query)
	if err != nil {
		logrus.Errorf("Error binding querystring data: %v", err)
	}
	streamer := service.CreateTweeterStreamer()
	stream, err := streamer.GetTweetStream(query.Q)
	if err != nil {
		logrus.Errorf("Error trying to stream from twitter: %v", err)
	}
	ctx.Stream(func(w io.Writer) bool {
		for msg := range stream.Messages {
			tweet := msg.(*twitter.Tweet).Text
			ctx.SSEvent("message", ResponseTweet{
				Tweet: tweet,
			})
			return true
		}
		return false
	})
}

// Handle tweets sample requests
// @Summary Get a stream of Server Sent Events containing tweets given a search query string
// @Router /tweets/sample [get]
// @Param q query string false "name search by q"
// @Produce json
// @Success 200 {array} handler.ResponseTweet
// @Failure 502 "Reached the maximum amount of simultaneous connections"
// @Tags tweets
func (h *Handler) handleTweetSample(ctx *gin.Context) {
	var query Query
	err := ctx.BindQuery(&query)
	if err != nil {
		logrus.Errorf("Error binding querystring data: %v", err)
	}
	streamer := service.CreateTweeterStreamer()
	sample, err := streamer.GetTweetSample(query.Q)
	if err != nil {
		logrus.Errorf("Error trying to get tweets sample: %v", err)
	}
	var tweets []ResponseTweet
	for _, t := range sample {
		tweets = append(tweets, ResponseTweet{Tweet: t.Text})
	}
	ctx.JSON(200, tweets)
}
