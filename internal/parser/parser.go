package parser

import (
	"lol_stats/internal/model"
)

func ParseUsername(username string) string {
	parsedUsername := ""

	for _, char := range username {
		if char == ' ' {
			parsedUsername += "%20"
		} else {
			parsedUsername += string(char)
		}
	}

	return parsedUsername + "/"
}

func ParseMatch(match model.Match, puuid string) model.GameStats {
	game := model.GameStats{}

	for _, participant := range match.Info.Participants {
		if participant.PUUID == puuid {
			game.Participant = participant
			game.GameDuration = uint32(match.Info.GameDuration)
			game.Win = participant.Win
		}
	}

	for _, opponent := range match.Info.Participants {
		if opponent.Role == game.Participant.Role {
			game.Opponent = opponent
		}
	}

	return game
}

func ParseMatchesInfo(matches []model.Match, account model.Account) []model.GameStats {
	games := []model.GameStats{}
	for _, match := range matches {
		game := ParseMatch(match, account.PUUID)
		games = append(games, game)
	}

	return games
}
