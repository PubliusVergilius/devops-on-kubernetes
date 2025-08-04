package main

import (
	"api/api"
	"api/collections"
	"api/lib"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var delay_startup = true
var fail_liveness = true
var fail_readiness = true

func init() {
	log.Println("Initializing app...")
	defaultPort := os.Getenv("API_PORT")
	if defaultPort == "" {
		os.Setenv("API_PORT", "8080")
	}
	lib.DotenvConfig(".env")
}

func main() {
	collections.ColorCollection()
	// Start Gin server
	port := os.Getenv("API_PORT")
	delay_startup, err := strconv.ParseBool(os.Getenv("DELAY_STARTUP"))
	fail_liveness, err := strconv.ParseBool(os.Getenv("FAIL_LIVENESS"))
	fail_readiness, err := strconv.ParseBool(os.Getenv("FAIL_READINESS"))
	if err != nil {
		log.Printf("Fail to parse boolean value from environment variables: %q", err)
	}
	fail_readiness = fail_readiness == true && rand.Float64() < 0.5
	log.Printf("Delay startup: %t", delay_startup)
	log.Printf("Fail liveness: %t", fail_liveness)
	log.Printf("Fail readiness: %t", fail_readiness)

	// Verify if port range or format is correct
	_, err = strconv.ParseUint(port, 10, 16)
	if err != nil {
		log.Fatalf("Conversion error or invalid port number: %q", err)
	}

	p := api.Probes{
		Delay:     delay_startup,
		Liveness:  fail_liveness,
		Readiness: fail_readiness,
	}

	if delay_startup {
		time.Sleep(60 * time.Second)
	}
	api.Router(port, p)
}
