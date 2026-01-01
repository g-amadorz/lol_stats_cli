package stats

import (
	"fmt"
	"lol_stats/internal/model"
	"math/rand"
)

func calculateScoreTop(model.Participant) float64 {

	return 10
}

func calculateScoreJungle(model.Participant) float64 {
	return float64(rand.Intn(10))
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
	fmt.Println(player.ChampionName)
	switch player.Lane {
	case "TOP":
		return calculateScoreTop(player)
	case "JUNGLE":
		return calculateScoreJungle(player)

	case "MID":
		return calculateScoreMid(player)

	case "BOT":
		return calculateScoreBot(player)

	case "SUPPORT":
		return calculateScoreSupport(player)

	default:
		return 0
	}

}
