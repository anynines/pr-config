package main

import (
	"log"

	"github.com/anynines/pr-config/web"
)

func main() {
	username, password, err := web.Credentials()
	if err != nil {
		log.Fatalf("You need to set HTTP_USERNAME and HTTP_PASSWORD.")
	}

	web.Run("4455", username, password)
}
