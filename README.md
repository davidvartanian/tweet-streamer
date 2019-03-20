# Tweet Streamer
This microservice provides a Server Sent Event endpoint to get a live stream of tweets based on a query string.

## Requirements
* docker
* docker-compose
* Twitter API credentials (see https://developer.twitter.com)

## Usage
* Setup environment variables (see .env.dist file)
* `docker-compose up -d --build`
* Follow logs: `docker-compose logs -f`
* Simultaneous connections to `/tweets/*` endpoints will be rejected according to MAX_CONCURRENCY environment variable (`/stats` won't be affected)

## Endpoints
### Tweets Stream
* GET /tweets/stream?q=some+keyword
* Produces a stream of tweets given a search query string `q`
* There is no response for this endpoint. Data is returned as Server-Sent-Events

### Tweets Sample
* GET /tweets/sample?q=some+keyword
* Produces a json list of tweets given a search query string `q`

### Stats
* GET /stats
* Produces a json object with the current status of the microservice

### Swagger Documentation
* GET /swagger/index.html
* Web page with information about the microservice and endpoints

## TODO
* Create unit and functional tests (see https://github.com/h2non/gock, http://onsi.github.io/ginkgo/, http://onsi.github.io/gomega/)
* Implement persistence of client requests
* Add request statistics of requests and gathered data

## See
* https://github.com/dghubble/go-twitter
