package api

import (
	"encoding/json"
	"fmt"
	"io"
	"lol_stats/internal/model"
	"lol_stats/internal/parser"
	"net/http"
)

const apiBasePuuidURL = "https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/"
const apiBaseMatchesURL = "https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/"
const apiBaseMatchURL = "https://americas.api.riotgames.com/lol/match/v5/matches/"
const apiPrefx = "?api_key="

func QueryAccount(username string, tagline string, apiKey string) (model.Account, error) {

	username = parser.ParseUsername(username)

	response, err := http.Get(apiBasePuuidURL + username + tagline + apiPrefx + apiKey)

	if err != nil {
		return model.Account{}, fmt.Errorf("unable to fetch Account: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return model.Account{}, fmt.Errorf("bad request: %d", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return model.Account{}, fmt.Errorf("unable to read response body %v", err)
	}

	account := model.Account{}

	if err := json.Unmarshal(responseBody, &account); err != nil {
		return model.Account{}, fmt.Errorf("unable to to unmarshal response body")
	}

	return account, nil
}

func QueryMatch(matchID string, apiKey string) (model.Match, error) {

	response, err := http.Get(apiBaseMatchURL + matchID + apiPrefx + apiKey)

	if response.StatusCode != http.StatusOK {
		return model.Match{}, fmt.Errorf("bad request for match query: %d", response.StatusCode)
	}

	if err != nil {
		return model.Match{}, fmt.Errorf("error")
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return model.Match{}, fmt.Errorf("error")
	}

	match := model.Match{}

	if err := json.Unmarshal(responseBody, &match); err != nil {
		return model.Match{}, fmt.Errorf("error")
	}

	return match, nil

}

// TODO: Implement a go routine to make the querying faster

func QueryMatches(account model.Account, apiKey string) ([]model.Match, error) {

	puuid := account.PUUID

	response, err := http.Get(apiBaseMatchesURL + puuid + "/" + "ids?type=ranked&start=0&count=20&api_key=" + apiKey)

	if response.StatusCode != http.StatusOK {
		return []model.Match{}, fmt.Errorf("bad request: %d", response.StatusCode)
	}

	if err != nil {
		return []model.Match{}, fmt.Errorf("unable to query matches: %w", err)
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return []model.Match{}, fmt.Errorf("unable to read response body: %w", err)
	}

	var matchIDs []string

	if err := json.Unmarshal(responseBody, &matchIDs); err != nil {
		return []model.Match{}, fmt.Errorf("unable to unmarshal matchIDs: %w", err)
	}

	matches := []model.Match{}

	for _, i := range matchIDs {
		match, err := QueryMatch(i, apiKey)

		if err != nil {
			return []model.Match{}, fmt.Errorf("unable to request match")
		}

		matches = append(matches, match)
	}

	return matches, nil

}
