package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"main/src/streamer/service"
)

// Handle tweets sample requests
// @Summary Get a stream of Server Sent Events containing tweets given a search query string
// @Router /tweets/sample [get]
// @Param q query string false "name search by q"
// @Produce json
// @Success 200 {array} handler.ResponseTweet
// @Failure 502 "Reached the maximum amount of simultaneous connections"
// @Tags tweets
func HandleTweetSample(ctx *gin.Context) {
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
