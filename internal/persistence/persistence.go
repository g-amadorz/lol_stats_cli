package persistence

import (
	"encoding/json"
	"log"
	"lol_stats/internal/api"
	"lol_stats/internal/model"
	"lol_stats/internal/stats"
	"os"
	"path/filepath"
)

// remove apiKey later
type Config struct {
	gamename string
	tag      string
	apiKey   string
}

type Performance struct {
	Idx         int
	Score       float64
	Participant model.Participant
}

const History = "history.json"

func GetConfigPath(file string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "lol_stats", file), nil
}

func saveConfig(data []byte, filename string) error {
	configPath, err := GetConfigPath(filename)
	if err != nil {
		return err
	}

	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func QueryPerformances(account model.Account, apiKey string) []Performance {

	matches, err := api.QueryMatches(account, apiKey)

	performances := []Performance{}

	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, match := range matches {

		if count >= 20 {
			break
		}

		for i, player := range match.Info.Participants {

			if player.Lane == "NONE" {
				break
			}

			if player.RiotIDGameName == account.GameName {

				performance := Performance{
					Idx:         i,
					Score:       stats.CalculateScore(player),
					Participant: player,
				}

				performances = append(performances, performance)

				count++

				break
			}
		}
	}

	return performances

}

func SaveGames(performances []Performance) error {

	data, err := json.Marshal(performances)

	if err != nil {
		log.Fatal(err)
	}

	return saveConfig(data, History)

}

func LoadGames() ([]Performance, error) {

	path, err := GetConfigPath(History)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(path)

	data := []Performance{}

	if err := json.Unmarshal(file, &data); err != nil {
		log.Fatal(err)
	}

	return data, nil

}
