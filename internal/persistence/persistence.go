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
const Account = "account.json"

func GetConfigPath(file string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "lol_stats", file), nil
}

func loadFile(filename string) ([]byte, error) {

	path, err := GetConfigPath(filename)

	if err != nil {
		log.Fatal(err)
	}

	return os.ReadFile(path)
}

func savePath(data []byte, filename string) error {
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

func SaveConfig(account model.Config) error {

	data, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}
	return savePath(data, Account)
}

func LoadConfig() (model.Config, error) {

	file, err := loadFile(Account)

	if err != nil {
		return model.Config{}, err
	}

	data := model.Config{}

	if err := json.Unmarshal(file, &data); err != nil {

		return model.Config{}, err
	}

	return data, nil

}

func QueryPerformances(config model.Config, apiKey string) []Performance {

	matches, err := api.QueryMatches(config.PUUID, apiKey)

	performances := []Performance{}

	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, match := range matches {

		if count >= 20 {
			break
		}

		for _, player := range match.Info.Participants {

			if player.Lane == "NONE" {
				break
			}

			if player.RiotIDGameName == config.Username {

				count++
				player.GameDuration = match.Info.GameDuration
				player.TotalMinionsKilled = player.LaneMinionsKilled + player.JungleMinionsKilled
				performance := Performance{
					Idx:         count,
					Score:       stats.CalculateScore(player),
					Participant: player,
				}

				performances = append(performances, performance)

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

	return savePath(data, History)

}

func LoadGames() ([]Performance, error) {

	file, err := loadFile(History)

	if err != nil {
		return nil, err
	}

	data := []Performance{}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data, nil

}
