package model

type MatchIDs struct {
	MatchIDs []string
}

type Match struct {
	Info MatchInfo `json:"info"`
}

type Account struct {
	PUUID string `json:"puuid"`

	GameName string `json:"gameName"`

	TagLine string `json:"tagLine"`
}

type MatchInfo struct {
	GameDuration int           `json:"gameDuration"` // in seconds
	GameMode     string        `json:"gameMode"`
	Participants []Participant `json:"participants"`
}

type Participant struct {
	RiotIDGameName string `json:"riotIdGameName"`
	RiotIDTagline  string `json:"riotIdTagline"`
	ChampionName   string `json:"championName"`
	ChampLevel     int    `json:"champLevel"`

	Kills   int  `json:"kills"`
	Deaths  int  `json:"deaths"`
	Assists int  `json:"assists"`
	Win     bool `json:"win"`

	TotalDamageDealtToChampions int `json:"totalDamageDealtToChampions"`
	TotalDamageTaken            int `json:"totalDamageTaken"`

	GoldEarned int `json:"goldEarned"`

	TimePlayed         int `json:"timePlayed"`
	TotalTimeSpentDead int `json:"totalTimeSpentDead"`

	TotalTimeCCDealt      int `json:"totalTimeCCDealt"`
	VisionScore           int `json:"visionScore"`
	TotalHealsOnTeammates int `json:"totalHealsOnTeammates"`
	TotalMinionsKilled    int `json:"totalMinionsKilled"`

	Lane         string `json:"lane"`
	Role         string `json:"role"`
	TeamPosition string `json:"teamPosition"`

	Placement int `json:"placement"`
}

type GameStats struct {
	ID               string
	Participant      Participant
	GameDuration     uint32
	PerformanceScore float64
	Opponent         Participant
	Win              bool
}
