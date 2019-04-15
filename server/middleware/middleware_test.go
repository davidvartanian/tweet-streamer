package middleware_test

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"main/server/middleware"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"
)

var (
	engine *gin.Engine
)

func GetServerMock() *gin.Engine {
	cs := middleware.GetConnectionStopper()
	cs.Setup(2)
	r := gin.Default()
	r.Use(cs.Limit())
	r.GET("/tweets/stream", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	return r
}

var _ = Describe("Middleware", func() {
	BeforeEach(func() {
		engine = GetServerMock()
		testServerListener, _ := net.Listen("tcp", ":9999")
		testServerSv := &http.Server{Handler: engine}
		go func() {
			logrus.Fatal(testServerSv.Serve(testServerListener))
		}()
	})

	Describe("Limiting concurrent connections", func() {
		Context("Up to 2", func() {
			It("should reject the third+ request", func() {
				var responses []int
				//var err error
				m := sync.Mutex{}
				wg := sync.WaitGroup{}
				for i := 0; i < 10; i++ {
					wg.Add(1)
					go func() {
						resp := sendRequest("GET", "http://localhost:9999/tweets/stream?q=test")
						m.Lock()
						responses = append(responses, resp.StatusCode)
						m.Unlock()
						wg.Done()
					}()
				}
				wg.Wait()
				sort.Ints(responses) // sort status codes to keep 503 error at last
				Expect(responses[len(responses)-1]).To(Equal(http.StatusServiceUnavailable))
				// kill server
				engine = nil
			})
		})
	})
})

var _ = AfterSuite(func() {
	//server = nil
})

func sendRequest(method string, url string) *http.Response {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequest(method, url, nil)
	resp, _ := httpClient.Do(req)
	defer resp.Body.Close()
	return resp
}
