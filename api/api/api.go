package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init () {
	log.Printf("Starting new Gin server...")
}

type Probes struct {
	Delay bool
	Liveness bool
	Readiness bool
}

func Router (port string, probes Probes) {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})	
	})

	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error on getting hostname: %q", err)
	}
	router.GET("/hostname", func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hostname": hostname,
	})	
})

	router.GET("/ready", func (ctx *gin.Context) {
		if (probes.Readiness) {
			ctx.AbortWithStatus(http.StatusServiceUnavailable)
		} else {
			ctx.Status(http.StatusOK)
		}
	})

	router.GET("/up", func (ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	router.GET("/healthy", func (ctx *gin.Context) {
		if (probes.Liveness) {
			ctx.AbortWithStatus(http.StatusServiceUnavailable)
		} else {
			ctx.Status(http.StatusOK)
		}

	})

	router.Run(fmt.Sprintf(":%s", port))
}

