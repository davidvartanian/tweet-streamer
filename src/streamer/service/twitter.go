package service

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *TwitterStreamer
)

type TwitterStreamer struct {
	client *twitter.Client
}

func CreateTweeterStreamer() *TwitterStreamer {
	once.Do(func() {
		instance = &TwitterStreamer{
			client: getTwitterClient(),
		}
	})
	return instance
}

func getTwitterClient() *twitter.Client {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TOKEN"), os.Getenv("TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

func getTwitterParams(q string) *twitter.StreamFilterParams {
	params := &twitter.StreamFilterParams{
		Track:         []string{q},
		StallWarnings: twitter.Bool(true),
	}
	return params
}

// Get a stream channel of tweets based on a query string
func (ts *TwitterStreamer) GetTweetStream(q string) (*twitter.Stream, error) {
	params := getTwitterParams(q)
	return ts.client.Streams.Filter(params)
}

// Get a sample of tweets based on a query string
func (ts *TwitterStreamer) GetTweetSample(q string) ([]twitter.Tweet, error) {
	search, _, err := ts.client.Search.Tweets(&twitter.SearchTweetParams{
		Query: q,
	})
	return search.Statuses, err
}
