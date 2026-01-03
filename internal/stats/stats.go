package stats

import (
	"lol_stats/internal/model"
	"math/rand"
)

func calculateScoreTop(model.Participant) float64 {

	return 10
}

func calculateScoreJungle(model.Participant) float64 {
	return float64(rand.Intn(10)) + 1
}

func calculateScoreMid(model.Participant) float64 {

	return 7
}

func calculateScoreBot(model.Participant) float64 {

	return 6
}
func calculateScoreSupport(model.Participant) float64 {

	return 3
}

func CalculateScore(player model.Participant) float64 {

	switch player.Lane {
	case "TOP":
		return calculateScoreTop(player)
	case "JUNGLE":
		return calculateScoreJungle(player)

	case "MIDDLE":
		return calculateScoreMid(player)

	case "BOTTOM":
		return calculateScoreBot(player)

	case "UTILITY":
		return calculateScoreSupport(player)

	default:
		return 0
	}

}
