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

func ParseMatch(match model.Match) {

}

func ParseMatches(matches []model.Match) {
	parsedMatches := []model.Match{}
	for _, match := range matches {
		parsedMatch := ParseMatch(match)
		parsedMatches = append(parsedMatches, parsedMatch)
	}
}
