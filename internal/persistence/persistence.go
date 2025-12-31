package persistence

import (
	"encoding/json"
	"log"
	"lol_stats/internal/api"
	"lol_stats/internal/model"
	"os"
	"path/filepath"
)

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

func QueryPerformances(account model.Account, apiKey string) []model.Participant {

	matches, err := api.QueryMatches(account, apiKey)

	performances := []model.Participant{}

	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		for _, player := range match.Info.Participants {
			if player.RiotIDGameName == account.GameName {
				performances = append(performances, player)
			}
		}
	}

	return performances

}

func SaveGames(performances []model.Participant) error {

	data, err := json.Marshal(performances)

	if err != nil {
		log.Fatal(err)
	}

	return saveConfig(data, History)

}

func LoadGames() ([]model.Participant, error) {

	path, err := GetConfigPath(History)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(path)

	data := []model.Participant{}

	if err := json.Unmarshal(file, &data); err != nil {
		log.Fatal(err)
	}

	return data, nil

}
