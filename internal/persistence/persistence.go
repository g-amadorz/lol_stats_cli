package persistence

import (
	"bufio"
	"encoding/json"
	"fmt"
	"lol_stats/internal/api"
	"lol_stats/internal/parser"
	"os"
	"path/filepath"
	"strings"
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
	username, _ := reader.ReadString('\n')

	fmt.Print("Tagline: ")
	tagLine, _ := reader.ReadString('\n')

	fmt.Print("Region (na1, euw1, kr, etc.): ")
	region, _ := reader.ReadString('\n')

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
		return nil
	}

	matchHistory, err := api.QueryMatches(account, apiKey)

	if err != nil {
		return nil
	}

	games := parser.ParseMatchesInfo(matchHistory, account)
}

func SaveHistory(overWrite bool) {

}

func ReadHistory() {

}
