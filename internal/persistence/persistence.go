package persistence

import (
	"bufio"
	"encoding/json"
	"fmt"
	"lol_stats/internal/api"
	"lol_stats/internal/model"
	"lol_stats/internal/parser"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	Username string `json:"gameName"`

	TagLine string `json:"tagLine"`

	Region string `json:"region"`
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".lol_stats", "config.json"), nil
}

func LoadOrCreateConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err == nil {
		var cfg Config
		if err := json.Unmarshal(data, &cfg); err == nil {
			return &cfg, nil
		}
	}

	fmt.Println("Welcome! Let's set up your League stats tracker.")
	fmt.Println("Config will be saved to:", configPath)
	fmt.Println()

	return InteractiveSetup()
}

func InteractiveSetup() (*Config, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	fmt.Print("Tagline: ")
	tagLine, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	fmt.Print("Region (na1, euw1, kr, etc.): ")
	region, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Username: strings.TrimSpace(username),
		TagLine:  strings.TrimSpace(tagLine),
		Region:   strings.TrimSpace(region),
	}

	if err := saveConfig(cfg); err != nil {
		return nil, err
	}

	fmt.Println("\n✓ Config saved! You're all set.")
	return cfg, nil
}

func saveConfig(cfg *Config) error {

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func LoadHistory(apiKey string) error {
	cfg, err := LoadOrCreateConfig()

	if err != nil {
		return err
	}

	username := cfg.Username
	tagLine := cfg.TagLine

	account, err := api.QueryAccount(username, tagLine, apiKey)

	if err != nil {
		return fmt.Errorf("failed to query account: %w", err)
	}

	matchHistory, err := api.QueryMatches(account, apiKey)

	if err != nil {
		return fmt.Errorf("failed to query match history: %w", err)
	}

	games := parser.ParseMatchesInfo(matchHistory, account)

	// will break if used have to add some bool to control or check if history.json exists already

	if err := ArchiveHistory(); err != nil && os.IsNotExist(err) {
		return fmt.Errorf("failed to archive history: %w", err)
	}

	file, err := os.Create("history.json")

	if err != nil {
		return fmt.Errorf("failed to create save file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(games); err != nil {
		return err
	}

	return nil
}

func ArchiveHistory() error {
	if _, err := os.Stat("history.json"); os.IsNotExist(err) {
		return err
	}
	date := time.Now().Format("2006-01-02")

	if err := os.Rename("history.json", "history_"+date+".json"); err != nil {
		fmt.Println("Error saving history:", err)
		return err
	}

	return nil
}

func ReadHistory() ([]model.GameStats, error) {
	file, err := os.ReadFile("history.json")

	if err != nil {
		return nil, fmt.Errorf("error reading match history file: %w", err)
	}

	var games []model.GameStats

	if err := json.Unmarshal(file, &games); err != nil {
		return nil, fmt.Errorf("error unmarshaling the match history file: %w", err)
	}

	return games, nil
}
