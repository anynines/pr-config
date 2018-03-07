package web

import (
	"fmt"
	"os"
)

func Credentials() (authUsername string, authPassword string, err error) {
	authUsername = os.Getenv("HTTP_USERNAME")
	authPassword = os.Getenv("HTTP_PASSWORD")

	if len(authUsername) < 1 || len(authPassword) < 1 {
		err = fmt.Errorf("Could not find credentials")
	}

	return authUsername, authPassword, err
}
