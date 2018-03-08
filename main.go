package main

import (
	"log"
	"os"

	"github.com/anynines/pr-config/web"
)

func main() {
	username, password, err := web.Credentials()
	if err != nil {
		log.Fatalf("You need to set HTTP_USERNAME and HTTP_PASSWORD.")
	}

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "4455"
	}

	backendURL := os.Getenv("BACKEND_URL")
	if len(backendURL) < 1 {
		log.Fatalf("You need to set BACKEND_URL.")
	}

	web.Run(port, backendURL, username, password)
}
