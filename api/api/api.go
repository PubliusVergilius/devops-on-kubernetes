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

func Router (port string) {
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
	router.Run(fmt.Sprintf(":%s", port))}

