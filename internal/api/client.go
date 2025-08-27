package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"lol_stats/internal/parser"
	"net/http"
	"os"
)

const API_BASE_PUUID_URL = "https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/"
const API_BASE_MATCHES_URL = "https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/"
const API_BASE_MATCH_URL = "https://americas.api.riotgames.com/lol/match/v5/matches/"
const API_PREFIX = "?api_key="

func LoadAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file")
	}

	api_key := os.Getenv("API_KEY")

	if api_key == "" {
		log.Fatal("Unable to load API_KEY")
	}

	return api_key

}

func QueryPuuid(username string, tagline string) error {

	api_key := LoadAPIKey()

	username = parser.ParseUsername(username)

	response, err := http.Get(API_BASE_PUUID_URL + username + tagline + API_PREFIX + api_key)

	if err != nil {
		return fmt.Errorf("Unable to fetch PUUID: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad request: ")
	}

}

func QueryMatches(puuid string) {

}

func QueryMatch(matchId string) {

}
