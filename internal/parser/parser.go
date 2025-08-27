package parser

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

func ParsePuuid() {

}

func ParseMatch() {

}

func ParseMatches() {

}
