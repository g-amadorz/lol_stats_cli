package stats

import (
	"lol_stats/internal/model"
	"math"
	"testing"
)

func TestCalculateKDAScore(t *testing.T) {
	tests := []struct {
		name     string
		kills    int
		deaths   int
		assists  int
		expected float64
	}{
		{
			name:     "perfect KDA (no deaths)",
			kills:    10,
			deaths:   0,
			assists:  15,
			expected: 30.0, // (10+15)/1 * 5 = 125, capped at 30
		},
		{
			name:     "excellent KDA",
			kills:    10,
			deaths:   2,
			assists:  12,
			expected: 30.0, // (10+12)/2 * 5 = 55, capped at 30
		},
		{
			name:     "good KDA",
			kills:    5,
			deaths:   2,
			assists:  5,
			expected: 25.0, // (5+5)/2 * 5 = 25
		},
		{
			name:     "average KDA",
			kills:    3,
			deaths:   3,
			assists:  6,
			expected: 15.0, // (3+6)/3 * 5 = 15
		},
		{
			name:     "poor KDA",
			kills:    1,
			deaths:   5,
			assists:  4,
			expected: 5.0, // (1+4)/5 * 5 = 5
		},
		{
			name:     "zero KDA",
			kills:    0,
			deaths:   5,
			assists:  0,
			expected: 0.0, // (0+0)/5 * 5 = 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateKDAScore(tt.kills, tt.deaths, tt.assists)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("calculateKDAScore(%d, %d, %d) = %.2f, want %.2f",
					tt.kills, tt.deaths, tt.assists, result, tt.expected)
			}
		})
	}
}

func TestCalculateDamageScore(t *testing.T) {
	tests := []struct {
		name         string
		totalDamage  int
		gameDuration int
		expected     float64
	}{
		{
			name:         "excellent damage",
			totalDamage:  30000,
			gameDuration: 1800, // 30 minutes
			expected:     25.0, // 1000 DPM, capped at 25
		},
		{
			name:         "good damage",
			totalDamage:  18000,
			gameDuration: 1800, // 30 minutes
			expected:     15.0, // 600 DPM
		},
		{
			name:         "average damage",
			totalDamage:  9000,
			gameDuration: 1800, // 30 minutes
			expected:     7.5,  // 300 DPM
		},
		{
			name:         "low damage",
			totalDamage:  3000,
			gameDuration: 1800, // 30 minutes
			expected:     2.5,  // 100 DPM
		},
		{
			name:         "zero duration",
			totalDamage:  10000,
			gameDuration: 0,
			expected:     0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateDamageScore(tt.totalDamage, tt.gameDuration)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("calculateDamageScore(%d, %d) = %.2f, want %.2f",
					tt.totalDamage, tt.gameDuration, result, tt.expected)
			}
		})
	}
}

func TestCalculateVisionScorePoints(t *testing.T) {
	tests := []struct {
		name         string
		visionScore  int
		gameDuration int
		expected     float64
	}{
		{
			name:         "excellent vision",
			visionScore:  60,
			gameDuration: 1800, // 30 minutes, 2.0 VSPM
			expected:     15.0,
		},
		{
			name:         "good vision",
			visionScore:  45,
			gameDuration: 1800, // 30 minutes, 1.5 VSPM
			expected:     11.25,
		},
		{
			name:         "average vision",
			visionScore:  30,
			gameDuration: 1800, // 30 minutes, 1.0 VSPM
			expected:     7.5,
		},
		{
			name:         "low vision",
			visionScore:  15,
			gameDuration: 1800, // 30 minutes, 0.5 VSPM
			expected:     3.75,
		},
		{
			name:         "zero duration",
			visionScore:  50,
			gameDuration: 0,
			expected:     0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateVisionScorePoints(tt.visionScore, tt.gameDuration)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("calculateVisionScorePoints(%d, %d) = %.2f, want %.2f",
					tt.visionScore, tt.gameDuration, result, tt.expected)
			}
		})
	}
}

func TestCalculateObjectiveScore(t *testing.T) {
	tests := []struct {
		name        string
		participant model.Participant
		minExpected float64
		maxExpected float64
	}{
		{
			name: "high CS and CC",
			participant: model.Participant{
				TotalMinionsKilled:    240, // 8 CS/min for 30 min = 10 points
				TotalTimeCCDealt:      100, // 5 points
				TotalHealsOnTeammates: 0,
			},
			minExpected: 15.0,
			maxExpected: 15.0,
		},
		{
			name: "support with healing",
			participant: model.Participant{
				TotalMinionsKilled:    40,   // 0.5 points
				TotalTimeCCDealt:      150,  // 5 points (capped)
				TotalHealsOnTeammates: 5000, // 5 points
			},
			minExpected: 10.5,
			maxExpected: 10.5,
		},
		{
			name: "average performance",
			participant: model.Participant{
				TotalMinionsKilled:    120, // 5 points
				TotalTimeCCDealt:      50,  // 2.5 points
				TotalHealsOnTeammates: 0,
			},
			minExpected: 7.5,
			maxExpected: 7.5,
		},
		{
			name: "low performance",
			participant: model.Participant{
				TotalMinionsKilled:    20,
				TotalTimeCCDealt:      10,
				TotalHealsOnTeammates: 0,
			},
			minExpected: 0.0,
			maxExpected: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateObjectiveScore(tt.participant)
			if result < tt.minExpected || result > tt.maxExpected {
				t.Errorf("calculateObjectiveScore() = %.2f, want between %.2f and %.2f",
					result, tt.minExpected, tt.maxExpected)
			}
		})
	}
}

func TestCalculateEfficiencyScore(t *testing.T) {
	tests := []struct {
		name         string
		participant  model.Participant
		gameDuration int
		minExpected  float64
		maxExpected  float64
	}{
		{
			name: "excellent efficiency",
			participant: model.Participant{
				TotalTimeSpentDead: 60,    // 60 seconds dead in 30 min = 96.7% alive
				GoldEarned:         13500, // 450 GPM
			},
			gameDuration: 1800,
			minExpected:  9.5,
			maxExpected:  10.0,
		},
		{
			name: "good efficiency",
			participant: model.Participant{
				TotalTimeSpentDead: 180,   // 180 seconds dead = 90% alive
				GoldEarned:         12000, // 400 GPM
			},
			gameDuration: 1800,
			minExpected:  8.5,
			maxExpected:  9.5,
		},
		{
			name: "poor efficiency",
			participant: model.Participant{
				TotalTimeSpentDead: 600,  // 600 seconds dead = 66.7% alive
				GoldEarned:         6000, // 200 GPM
			},
			gameDuration: 1800,
			minExpected:  5.0,
			maxExpected:  6.5,
		},
		{
			name: "zero duration",
			participant: model.Participant{
				TotalTimeSpentDead: 0,
				GoldEarned:         10000,
			},
			gameDuration: 0,
			minExpected:  0.0,
			maxExpected:  0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateEfficiencyScore(tt.participant, tt.gameDuration)
			if result < tt.minExpected || result > tt.maxExpected {
				t.Errorf("calculateEfficiencyScore() = %.2f, want between %.2f and %.2f",
					result, tt.minExpected, tt.maxExpected)
			}
		})
	}
}

func TestCalculateMultikillBonus(t *testing.T) {
	tests := []struct {
		name        string
		participant model.Participant
		expected    float64
	}{
		{
			name: "pentakill",
			participant: model.Participant{
				PentaKills: 1,
			},
			expected: 8.0,
		},
		{
			name: "quadrakill",
			participant: model.Participant{
				QuadraKills: 1,
			},
			expected: 4.0,
		},
		{
			name: "triple kill",
			participant: model.Participant{
				TripleKills: 1,
			},
			expected: 2.0,
		},
		{
			name: "double kills",
			participant: model.Participant{
				DoubleKills: 3,
			},
			expected: 1.5,
		},
		{
			name: "mixed multikills",
			participant: model.Participant{
				DoubleKills: 2,
				TripleKills: 1,
				QuadraKills: 1,
			},
			expected: 7.0, // 1.0 + 2.0 + 4.0
		},
		{
			name: "excessive multikills (capped)",
			participant: model.Participant{
				DoubleKills: 10,
				TripleKills: 5,
			},
			expected: 10.0, // capped at 10
		},
		{
			name:        "no multikills",
			participant: model.Participant{},
			expected:    0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateMultikillBonus(tt.participant)
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("calculateMultikillBonus() = %.2f, want %.2f", result, tt.expected)
			}
		})
	}
}

func TestCalculatePerformanceScore(t *testing.T) {
	tests := []struct {
		name         string
		participant  model.Participant
		gameDuration int
		checkWin     bool
	}{
		{
			name: "excellent performance with win",
			participant: model.Participant{
				Kills:                       10,
				Deaths:                      2,
				Assists:                     15,
				Win:                         true,
				TotalDamageDealtToChampions: 30000,
				VisionScore:                 60,
				TotalMinionsKilled:          240,
				TotalTimeCCDealt:            100,
				TotalTimeSpentDead:          60,
				GoldEarned:                  13500,
				TripleKills:                 1,
			},
			gameDuration: 1800,
			checkWin:     true,
		},
		{
			name: "poor performance with loss",
			participant: model.Participant{
				Kills:                       1,
				Deaths:                      10,
				Assists:                     3,
				Win:                         false,
				TotalDamageDealtToChampions: 5000,
				VisionScore:                 15,
				TotalMinionsKilled:          80,
				TotalTimeCCDealt:            20,
				TotalTimeSpentDead:          600,
				GoldEarned:                  6000,
			},
			gameDuration: 1800,
			checkWin:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePerformanceScore(tt.participant, tt.gameDuration)

			// Verify total score is sum of components
			expectedTotal := result.KDAScore + result.DamageScore + result.VisionScore +
				result.ObjectiveScore + result.EfficiencyScore + result.MultikillBonus + result.WinBonus

			if math.Abs(result.TotalScore-expectedTotal) > 0.01 {
				t.Errorf("TotalScore = %.2f, but sum of components = %.2f",
					result.TotalScore, expectedTotal)
			}

			// Verify win bonus
			if tt.checkWin && result.WinBonus != 15.0 {
				t.Errorf("Expected WinBonus = 15.0 for win, got %.2f", result.WinBonus)
			}

			if !tt.checkWin && result.WinBonus != 0.0 {
				t.Errorf("Expected WinBonus = 0.0 for loss, got %.2f", result.WinBonus)
			}

			// Verify all scores are non-negative
			if result.KDAScore < 0 || result.DamageScore < 0 || result.VisionScore < 0 ||
				result.ObjectiveScore < 0 || result.EfficiencyScore < 0 || result.MultikillBonus < 0 {
				t.Error("All score components should be non-negative")
			}
		})
	}
}

func TestCalculateMatchScore(t *testing.T) {
	match := model.Match{
		Info: model.MatchInfo{
			GameDuration: 1800,
			Participants: []model.Participant{
				{
					PUUID:                       "player1",
					Kills:                       10,
					Deaths:                      2,
					Assists:                     15,
					Win:                         true,
					TotalDamageDealtToChampions: 30000,
					VisionScore:                 60,
					TotalMinionsKilled:          240,
					TotalTimeCCDealt:            100,
					TotalTimeSpentDead:          60,
					GoldEarned:                  13500,
				},
				{
					PUUID:                       "player2",
					Kills:                       5,
					Deaths:                      5,
					Assists:                     10,
					Win:                         true,
					TotalDamageDealtToChampions: 20000,
					VisionScore:                 40,
					TotalMinionsKilled:          180,
					TotalTimeCCDealt:            80,
					TotalTimeSpentDead:          120,
					GoldEarned:                  11000,
				},
			},
		},
	}

	scores := CalculateMatchScore(match)

	// Verify we have scores for both players
	if len(scores) != 2 {
		t.Fatalf("Expected 2 scores, got %d", len(scores))
	}

	// Verify scores exist for both players
	if _, exists := scores["player1"]; !exists {
		t.Error("Expected score for player1")
	}

	if _, exists := scores["player2"]; !exists {
		t.Error("Expected score for player2")
	}

	// Verify player1 has higher score (better performance)
	if scores["player1"].TotalScore <= scores["player2"].TotalScore {
		t.Errorf("Expected player1 score (%.2f) > player2 score (%.2f)",
			scores["player1"].TotalScore, scores["player2"].TotalScore)
	}
}

func TestGetAverageScore(t *testing.T) {
	testPUUID := "test-player"

	matches := []model.Match{
		{
			Info: model.MatchInfo{
				GameDuration: 1800,
				Participants: []model.Participant{
					{
						PUUID:                       testPUUID,
						Kills:                       10,
						Deaths:                      2,
						Assists:                     15,
						Win:                         true,
						TotalDamageDealtToChampions: 30000,
						VisionScore:                 60,
						TotalMinionsKilled:          240,
						TotalTimeCCDealt:            100,
						TotalTimeSpentDead:          60,
						GoldEarned:                  13500,
					},
				},
			},
		},
		{
			Info: model.MatchInfo{
				GameDuration: 1800,
				Participants: []model.Participant{
					{
						PUUID:                       testPUUID,
						Kills:                       5,
						Deaths:                      5,
						Assists:                     10,
						Win:                         false,
						TotalDamageDealtToChampions: 20000,
						VisionScore:                 40,
						TotalMinionsKilled:          180,
						TotalTimeCCDealt:            80,
						TotalTimeSpentDead:          120,
						GoldEarned:                  11000,
					},
				},
			},
		},
	}

	result := GetAverageScore(matches, testPUUID)

	// Verify average is calculated
	if result.TotalScore == 0 {
		t.Error("Expected non-zero average score")
	}

	// Calculate expected scores manually
	score1 := CalculatePerformanceScore(matches[0].Info.Participants[0], 1800)
	score2 := CalculatePerformanceScore(matches[1].Info.Participants[0], 1800)

	expectedTotal := (score1.TotalScore + score2.TotalScore) / 2
	if math.Abs(result.TotalScore-expectedTotal) > 0.01 {
		t.Errorf("Expected average total score %.2f, got %.2f", expectedTotal, result.TotalScore)
	}
}

func TestGetAverageScoreEmptyMatches(t *testing.T) {
	result := GetAverageScore([]model.Match{}, "test-player")

	// Verify all scores are zero for empty matches
	if result.TotalScore != 0 || result.KDAScore != 0 || result.DamageScore != 0 {
		t.Error("Expected all scores to be zero for empty matches")
	}
}

func TestGetAverageScorePlayerNotFound(t *testing.T) {
	matches := []model.Match{
		{
			Info: model.MatchInfo{
				GameDuration: 1800,
				Participants: []model.Participant{
					{
						PUUID: "other-player",
						Kills: 10,
					},
				},
			},
		},
	}

	result := GetAverageScore(matches, "non-existent-player")

	// Verify all scores are zero when player not found
	if result.TotalScore != 0 {
		t.Error("Expected zero score when player not found in matches")
	}
}
