package api

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"lol_stats/internal/model"
	"lol_stats/internal/parser"
	"net/http"
	"os"
)

const apiBasePuuidURL = "https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/"
const apiBaseMatchesURL = "https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/"
const apiBaseMatchURL = "https://americas.api.riotgames.com/lol/match/v5/matches/"
const apiPrefx = "?api_key="

type Match model.Match

type Account struct {
	PUUID string `json:"puuid"`

	GameName string `json:"gameName"`

	TagLine string `json:"tagLine"`
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

	response, err := http.Get(apiBasePuuidURL + username + tagline + apiPrefx + apiKey)

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

	account := Account{}

	if err := json.Unmarshal(responseBody, &account); err != nil {
		return Account{}, fmt.Errorf("unable to to unmarshal response body")
	}

	return account, nil
}

func QueryMatch(matchID string) (Match, error) {
	apiKey := LoadAPIKey()

	response, err := http.Get(apiBaseMatchURL + matchID + apiPrefx + apiKey)

	if response.StatusCode != http.StatusOK {
		return Match{}, fmt.Errorf("bad request for match query: %d", response.StatusCode)
	}

	if err != nil {
		return Match{}, fmt.Errorf("error")
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return Match{}, fmt.Errorf("error")
	}

	match := Match{}

	if err := json.Unmarshal(responseBody, &match); err != nil {
		return Match{}, fmt.Errorf("error")
	}

	return match, nil

}

func QueryMatches(account Account) ([]Match, error) {
	apiKey := LoadAPIKey()

	puuid := account.PUUID

	response, err := http.Get(apiBaseMatchesURL + puuid + "/" + "ids?start=0&count=20&api_key=" + apiKey)

	if response.StatusCode != http.StatusOK {
		return []Match{}, fmt.Errorf("bad request: %d", response.StatusCode)
	}

	if err != nil {
		return []Match{}, fmt.Errorf("unable to query matches: %w", err)
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return []Match{}, fmt.Errorf("unable to read response body: %w", err)
	}

	var matchIDs []string

	if err := json.Unmarshal(responseBody, &matchIDs); err != nil {
		return []Match{}, fmt.Errorf("unable to unmarshal matchIDs: %w", err)
	}

	matches := []Match{}

	for _, i := range matchIDs {
		match, err := QueryMatch(i)

		if err != nil {
			return []Match{}, fmt.Errorf("unable to request match")
		}

		matches = append(matches, match)
	}

	return matches, nil

}
