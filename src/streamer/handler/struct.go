package handler

type Query struct {
	Q string `form:"q"`
}

type ResponseTweet struct {
	Tweet string `json:"tweet"`
}
