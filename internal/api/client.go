package api

import (
	"encoding/json"
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

type Account struct {
	PUUID string `json:"puuid"`

	GameName string `json:"gameName"`

	TagLine string `json:"tagLine"`
}

type Match struct {
}

func LoadAPIKey() string {
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

func QueryAccount(username string, tagline string) (Account, error) {

	apiKey := LoadAPIKey()

	username = parser.ParseUsername(username)

	response, err := http.Get(API_BASE_PUUID_URL + username + tagline + API_PREFIX + apiKey)

	if err != nil {
		return Account{}, fmt.Errorf("unable to fetch Account: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return Account{}, fmt.Errorf("bad request: %d", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return Account{}, fmt.Errorf("unable to read response body %v", err)
	}

	account := &Account{}

	if err := json.Unmarshal(responseBody, account); err != nil {
		return Account{}, fmt.Errorf("unable to to unmarshal response body")
	}

	return *account, nil
}

func QueryMatches(account Account) ([]Match, error) {
	apiKey := LoadAPIKey()

	puuid := account.PUUID

	response, err := http.Get(API_BASE_MATCHES_URL + puuid + apiKey)

	if err != nil {
		return []Match{}, fmt.Errorf("error")
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return []Match{}, fmt.Errorf("error")
	}

}

func QueryMatch(matchID string) {

}
