package stats

import (
	"lol_stats/internal/model"
	"math"
)

// Non-Support Distribution = [KDA : 25, Damage : 25, Vision : 10, Farm : 25, Opponent: 15]

type Role string

const (
	RoleTop     Role = "TOP"
	RoleJungle  Role = "JUNGLE"
	RoleMid     Role = "MID"
	RoleADC     Role = "ADC"
	RoleSupport Role = "SUPPORT"
)

// PerformanceScore represents the calculated score for a player's performance
type PerformanceScore struct {
	TotalScore      float64
	KDAScore        float64
	DamageScore     float64
	VisionScore     float64
	ObjectiveScore  float64
	EfficiencyScore float64
	MultikillBonus  float64
	WinBonus        float64
}

func CalculatePerformanceScore(p model.Participant, gameDuration int) PerformanceScore {
	score := PerformanceScore{}

	score.KDAScore = calculateKDAScore(p.Kills, p.Deaths, p.Assists)

	score.DamageScore = calculateDamageScore(p.TotalDamageDealtToChampions, gameDuration)

	score.VisionScore = calculateVisionScorePoints(p.VisionScore, gameDuration)

	score.ObjectiveScore = calculateObjectiveScore(p)

	score.EfficiencyScore = calculateEfficiencyScore(p, gameDuration)

	score.MultikillBonus = calculateMultikillBonus(p)

	if p.Win {
		score.WinBonus = 15.0
	}

	score.TotalScore = score.KDAScore + score.DamageScore + score.VisionScore +
		score.ObjectiveScore + score.EfficiencyScore + score.MultikillBonus + score.WinBonus

	return score
}

func calculateKDAScore(kills, deaths, assists int) float64 {
	// Avoid division by zero
	if deaths == 0 {
		deaths = 1
	}

	kda := float64(kills+assists) / float64(deaths)

	// Scale KDA to 0-30 points
	// Perfect KDA (10+) = 30 points
	// Good KDA (5) = 20 points
	// Average KDA (3) = 15 points
	// Poor KDA (1) = 5 points
	score := math.Min(kda*5, 30)

	return score
}

func calculateDamageScore(totalDamage, gameDuration int) float64 {
	if gameDuration == 0 {
		return 0
	}

	dpm := float64(totalDamage) / (float64(gameDuration) / 60.0)

	// Scale DPM to 0-25 points
	// 1000+ DPM = 25 points (excellent)
	// 600 DPM = 15 points (good)
	// 300 DPM = 7.5 points (average)
	score := math.Min((dpm/1000)*25, 25)

	return score
}

func calculateVisionScorePoints(visionScore, gameDuration int) float64 {
	if gameDuration == 0 {
		return 0
	}

	vspm := float64(visionScore) / (float64(gameDuration) / 60.0)

	// Scale vision score to 0-15 points
	// 2.0+ VSPM = 15 points (excellent)
	// 1.5 VSPM = 11.25 points (good)
	// 1.0 VSPM = 7.5 points (average)
	score := math.Min((vspm/2.0)*15, 15)

	return score
}

// calculateObjectiveScore calculates score based on CS and support contributions
func calculateObjectiveScore(p model.Participant) float64 {
	score := 0.0

	// CS contribution (0-10 points)
	// 8+ CS/min = 10 points
	csScore := math.Min(float64(p.TotalMinionsKilled)/80, 10)
	score += csScore

	// Healing contribution (0-5 points) - mainly for supports
	if p.TotalHealsOnTeammates > 0 {
		healScore := math.Min(float64(p.TotalHealsOnTeammates)/5000, 5)
		score += healScore
	}

	// CC contribution (0-5 points)
	ccScore := math.Min(float64(p.TotalTimeCCDealt)/100, 5)
	score += ccScore

	return math.Min(score, 15)
}

// calculateEfficiencyScore calculates score based on death time and gold efficiency
func calculateEfficiencyScore(p model.Participant, gameDuration int) float64 {
	if gameDuration == 0 {
		return 0
	}

	score := 0.0

	// Alive time percentage (0-5 points)
	aliveTimePercentage := 1.0 - (float64(p.TotalTimeSpentDead) / float64(gameDuration))
	score += aliveTimePercentage * 5

	// Gold per minute (0-5 points)
	gpm := float64(p.GoldEarned) / (float64(gameDuration) / 60.0)
	// 450+ GPM = 5 points
	gpmScore := math.Min((gpm/450)*5, 5)
	score += gpmScore

	return math.Min(score, 10)
}

func calculateMultikillBonus(p model.Participant) float64 {
	bonus := 0.0

	bonus += float64(p.DoubleKills) * 0.5
	bonus += float64(p.TripleKills) * 2.0
	bonus += float64(p.QuadraKills) * 4.0
	bonus += float64(p.PentaKills) * 8.0

	return math.Min(bonus, 10)
}

// CalculateMatchScore calculates scores for all participants in a match
func CalculateMatchScore(match model.Match) map[string]PerformanceScore {
	scores := make(map[string]PerformanceScore)

	for _, participant := range match.Info.Participants {
		score := CalculatePerformanceScore(participant, match.Info.GameDuration)
		scores[participant.PUUID] = score
	}

	return scores
}

func GetAverageScore(matches []model.Match, puuid string) PerformanceScore {
	if len(matches) == 0 {
		return PerformanceScore{}
	}

	var total PerformanceScore
	count := 0

	for _, match := range matches {
		for _, p := range match.Info.Participants {
			if p.PUUID == puuid {
				score := CalculatePerformanceScore(p, match.Info.GameDuration)
				total.TotalScore += score.TotalScore
				total.KDAScore += score.KDAScore
				total.DamageScore += score.DamageScore
				total.VisionScore += score.VisionScore
				total.ObjectiveScore += score.ObjectiveScore
				total.EfficiencyScore += score.EfficiencyScore
				total.MultikillBonus += score.MultikillBonus
				total.WinBonus += score.WinBonus
				count++
			}
		}
	}

	if count == 0 {
		return PerformanceScore{}
	}

	divisor := float64(count)
	return PerformanceScore{
		TotalScore:      total.TotalScore / divisor,
		KDAScore:        total.KDAScore / divisor,
		DamageScore:     total.DamageScore / divisor,
		VisionScore:     total.VisionScore / divisor,
		ObjectiveScore:  total.ObjectiveScore / divisor,
		EfficiencyScore: total.EfficiencyScore / divisor,
		MultikillBonus:  total.MultikillBonus / divisor,
		WinBonus:        total.WinBonus / divisor,
	}
}
