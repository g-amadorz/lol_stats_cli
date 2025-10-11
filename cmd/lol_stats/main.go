package main

import (
	"fmt"
	"log"
	"lol_stats/internal/api"
	"lol_stats/internal/parser"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	apiKey := loadAPIKey()
	account, err := api.QueryAccount("FREE PALESTINE", "tox", apiKey)
	if err != nil {
		log.Fatal(err)
	}
	matches, err := api.QueryMatches(account, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	games := parser.ParseMatchesInfo(matches, account)

	fmt.Println(games)
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
