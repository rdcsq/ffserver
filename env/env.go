package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ListeningAddress string
	AuthSecret       []byte
	Domain           string
)

func LoadEnv() {
	// ignore error
	godotenv.Load()

	ListeningAddress = os.Getenv("LISTENING_ADDRESS")
	if ListeningAddress == "" {
		ListeningAddress = "0.0.0.0:3000"
	}

	if os.Getenv("AUTH_SECRET") == "" {
		log.Println("You need to set an auth secret.")
		os.Exit(1)
	}

	AuthSecret = []byte(os.Getenv("AUTH_SECRET"))

	Domain = os.Getenv("DOMAIN")
	if Domain == "" {
		Domain = "http://0.0.0.0:3000"
	}
}
