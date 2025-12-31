package persistence_test

import (
	"encoding/json"
	"log"
	"lol_stats/internal/model"
	"lol_stats/internal/persistence"
	"os"
	"slices"
	"testing"
)

func TestSaveGames(t *testing.T) {
	tests := []struct {
		name         string
		performances []model.Participant
	}{
		{
			name: "saves performances to history.json",
			performances: []model.Participant{
				{
					RiotIDGameName: "TestPlayer",
					ChampionName:   "Ahri",
					Kills:          10,
					Deaths:         2,
					Assists:        5,
				},
			},
		},
		{
			name:         "saves empty performances list",
			performances: []model.Participant{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up test file
			configPath, err := persistence.GetConfigPath("history.json")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(configPath)

			persistence.SaveGames(tt.performances)

			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				t.Errorf("history.json was not created at %s", configPath)
			}

			data, err := os.ReadFile(configPath)
			if err != nil {
				t.Errorf("Failed to read history.json: %v", err)
			}

			var loaded []model.Participant
			if err := json.Unmarshal(data, &loaded); err != nil {
				t.Errorf("Invalid JSON in history.json: %v", err)
			}

			if len(loaded) != len(tt.performances) {
				t.Errorf("Expected %d performances, got %d", len(tt.performances), len(loaded))
			}
		})
	}
}

func TestLoadGames(t *testing.T) {
	tests := []struct {
		name         string
		performances []model.Participant
		want         []model.Participant
	}{
		{
			name: "Normal read from file",

			performances: []model.Participant{
				{
					RiotIDGameName: "TestPlayer",
					ChampionName:   "Ahri",
					Kills:          10,
					Deaths:         2,
					Assists:        5,
				},

				{
					RiotIDGameName: "TestPlayer",
					ChampionName:   "Ammumu",
					Kills:          3,
					Deaths:         2,
					Assists:        5,
				},
			},

			want: []model.Participant{
				{
					RiotIDGameName: "TestPlayer",
					ChampionName:   "Ahri",
					Kills:          10,
					Deaths:         2,
					Assists:        5,
				},

				{
					RiotIDGameName: "TestPlayer",
					ChampionName:   "Ammumu",
					Kills:          3,
					Deaths:         2,
					Assists:        5,
				},
			},
		},
		{
			name:         "empty",
			performances: []model.Participant{},
			want:         []model.Participant{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPath, err := persistence.GetConfigPath("history.json")

			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(configPath)

			persistence.SaveGames(tt.performances)

			matches, err := persistence.LoadGames()

			if err != nil {
				log.Fatal(err)
			}

			if !slices.Equal(matches, tt.performances) {
				log.Fatal("Loaded matches don't match expectation")
			}

		})
	}

}
