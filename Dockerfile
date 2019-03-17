FROM golang:latest
WORKDIR /app
COPY . /app
RUN go mod download
RUN swag init -g server/server.go
RUN go build -o tweet_streamer
EXPOSE $SERVICE_PORT