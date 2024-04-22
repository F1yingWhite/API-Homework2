// package main

// import (
// 	"api2/config"
// 	"api2/models"
// 	"api2/server"
// 	"api2/server/middlewares"
// 	"fmt"
// 	"log"
// 	"sync/atomic"
// 	"time"
// )

// func Init() {
// 	cfg, err := config.ReadConfig()
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	config.InitLog(cfg)
// 	models.InitDB(cfg)

// 	go func() {
// 		for {
// 			time.Sleep(time.Second)
// 			requestsLastSecond := atomic.SwapInt64(&middlewares.RequestCounter, 0)
// 			fmt.Printf("\rRequests in last second: %d", requestsLastSecond)
// 		}
// 	}()

// }

// func main() {
// 	Init()
// 	api := server.InitRouter()

// 	err := api.Run(":8888")
// 	if err != nil {
// 		log.Panicln(err)
// 	}
// }

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "status"},
	)
)

func init() {
	// Register custom metrics with Prometheus.
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {
	// Create a new Gin router.
	router := gin.Default()

	// Register Prometheus metrics handler.
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Define a route handler.
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})

		// Increment the HTTP request counter.
		httpRequestsTotal.WithLabelValues("GET", "200").Inc()
	})

	// Run the HTTP server.
	router.Run(":8080")
}
