package stats

import (
	"lol_stats/internal/model"
)

func calculateScoreTop(p model.Participant) float64 {
	score := 0.0

	// KDA component (0-30 points)
	kda := 0.0
	if p.Deaths > 0 {
		kda = float64(p.Kills+p.Assists) / float64(p.Deaths)
	} else {
		kda = float64(p.Kills + p.Assists)
	}
	score += min(kda*3, 30.0)

	// CS component (0-25 points) - 7+ CS/min is excellent for top
	if p.TimePlayed > 0 {
		csPerMin := float64(p.TotalMinionsKilled) / (float64(p.TimePlayed) / 60.0)
		score += min(csPerMin*3, 25.0)
	}

	// Damage dealt (0-20 points) - normalize to 30k = 20 points
	score += min(float64(p.TotalDamageDealtToChampions)/1500.0, 20.0)

	// Vision score (0-10 points) - 1.5 vision/min is good
	if p.TimePlayed > 0 {
		visionPerMin := float64(p.VisionScore) / (float64(p.TimePlayed) / 60.0)
		score += min(visionPerMin*5, 10.0)
	}

	// Win bonus (0-15 points)
	if p.Win {
		score += 15.0
	}

	return score
}

func calculateScoreJungle(p model.Participant) float64 {
	score := 0.0

	// KDA component (0-25 points)
	kda := 0.0
	if p.Deaths > 0 {
		kda = float64(p.Kills+p.Assists) / float64(p.Deaths)
	} else {
		kda = float64(p.Kills + p.Assists)
	}
	score += min(kda*2.5, 25.0)

	// Kill participation is crucial for junglers
	// Estimate: (K+A) should be high relative to team total
	// Give bonus for high K+A
	killParticipation := float64(p.Kills + p.Assists)
	score += min(killParticipation*0.8, 20.0)

	// Vision score (0-15 points) - junglers need high vision
	if p.TimePlayed > 0 {
		visionPerMin := float64(p.VisionScore) / (float64(p.TimePlayed) / 60.0)
		score += min(visionPerMin*7, 15.0)
	}

	// Damage dealt (0-15 points)
	score += min(float64(p.TotalDamageDealtToChampions)/1800.0, 15.0)

	// CC duration (0-10 points) - junglers provide CC for ganks
	if p.TimePlayed > 0 {
		ccPerMin := (float64(p.TotalTimeCCDealt) / 1000.0) / (float64(p.TimePlayed) / 60.0)
		score += min(ccPerMin*2, 10.0)
	}

	// Win bonus
	if p.Win {
		score += 15.0
	}

	return score
}

func calculateScoreMid(p model.Participant) float64 {
	score := 0.0

	// KDA component (0-30 points) - carries need good KDA
	kda := 0.0
	if p.Deaths > 0 {
		kda = float64(p.Kills+p.Assists) / float64(p.Deaths)
	} else {
		kda = float64(p.Kills + p.Assists)
	}
	score += min(kda*3, 30.0)

	// Damage dealt (0-30 points) - mid laners are primary damage dealers
	score += min(float64(p.TotalDamageDealtToChampions)/1200.0, 30.0)

	// CS component (0-20 points) - 7+ CS/min is good
	if p.TimePlayed > 0 {
		csPerMin := float64(p.TotalMinionsKilled) / (float64(p.TimePlayed) / 60.0)
		score += min(csPerMin*2.5, 20.0)
	}

	// Vision (0-5 points) - less critical but still matters
	if p.TimePlayed > 0 {
		visionPerMin := float64(p.VisionScore) / (float64(p.TimePlayed) / 60.0)
		score += min(visionPerMin*3, 5.0)
	}

	// Win bonus
	if p.Win {
		score += 15.0
	}

	return score
}

func calculateScoreBot(p model.Participant) float64 {
	score := 0.0

	// KDA component (0-30 points)
	kda := 0.0
	if p.Deaths > 0 {
		kda = float64(p.Kills+p.Assists) / float64(p.Deaths)
	} else {
		kda = float64(p.Kills + p.Assists)
	}
	score += min(kda*3, 30.0)

	// Damage dealt (0-30 points) - ADCs are primary damage dealers
	score += min(float64(p.TotalDamageDealtToChampions)/1200.0, 30.0)

	// CS component (0-25 points) - CS is critical for ADCs
	if p.TimePlayed > 0 {
		csPerMin := float64(p.TotalMinionsKilled) / (float64(p.TimePlayed) / 60.0)
		score += min(csPerMin*3, 25.0)
	}

	// Vision (0-5 points)
	if p.TimePlayed > 0 {
		visionPerMin := float64(p.VisionScore) / (float64(p.TimePlayed) / 60.0)
		score += min(visionPerMin*3, 5.0)
	}

	// Win bonus
	if p.Win {
		score += 10.0
	}

	return score
}

func calculateScoreSupport(p model.Participant) float64 {
	score := 0.0

	// KDA component (0-20 points) - assists matter most
	kda := 0.0
	if p.Deaths > 0 {
		kda = float64(p.Kills+p.Assists) / float64(p.Deaths)
	} else {
		kda = float64(p.Kills + p.Assists)
	}
	score += min(kda*2, 20.0)

	// Assist count directly (0-20 points)
	score += min(float64(p.Assists)*1.5, 20.0)

	// Vision score (0-30 points) - CRITICAL for supports
	if p.TimePlayed > 0 {
		visionPerMin := float64(p.VisionScore) / (float64(p.TimePlayed) / 60.0)
		score += min(visionPerMin*10, 30.0)
	}

	// CC dealt (0-15 points) - supports provide CC
	if p.TimePlayed > 0 {
		ccPerMin := (float64(p.TotalTimeCCDealt) / 1000.0) / (float64(p.TimePlayed) / 60.0)
		score += min(ccPerMin*3, 15.0)
	}

	// Healing (0-10 points) - if applicable
	if p.TotalHealsOnTeammates > 0 {
		score += min(float64(p.TotalHealsOnTeammates)/1000.0, 10.0)
	}

	// Win bonus
	if p.Win {
		score += 15.0
	}

	return score
}

// Helper function for min
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
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
