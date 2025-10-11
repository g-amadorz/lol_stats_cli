package main

import (
	"log"
	"lol_stats/cmd"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	apiKey := loadAPIKey()
	cmd.Execute(apiKey)
}

func loadAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file")
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("Unable to load API_KEY")
	}

	return apiKey

}
