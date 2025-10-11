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
	// Player Identity
	PUUID          string `json:"puuid"`
	RiotIDGameName string `json:"riotIdGameName"`
	RiotIDTagline  string `json:"riotIdTagline"`
	ChampionName   string `json:"championName"`
	ChampionID     int    `json:"championId"`
	ChampLevel     int    `json:"champLevel"`

	// Core Performance Stats
	Kills   int  `json:"kills"`
	Deaths  int  `json:"deaths"`
	Assists int  `json:"assists"`
	Win     bool `json:"win"`

	// Damage Stats
	TotalDamageDealtToChampions int `json:"totalDamageDealtToChampions"`
	TotalDamageTaken            int `json:"totalDamageTaken"`

	// Economy
	GoldEarned int `json:"goldEarned"`

	// Combat Details
	DoubleKills   int `json:"doubleKills"`
	TripleKills   int `json:"tripleKills"`
	QuadraKills   int `json:"quadraKills"`
	PentaKills    int `json:"pentaKills"`
	KillingSprees int `json:"killingSprees"`

	// Game Participation
	TimePlayed         int `json:"timePlayed"`
	TotalTimeSpentDead int `json:"totalTimeSpentDead"`

	// Outside Scores
	TotalTimeCCDealt      int `json:"totalTimeCCDealt"`
	VisionScore           int `json:"visionScore"`
	TotalHealsOnTeammates int `json:"totalHealsOnTeammates"`
	TotalMinionsKilled    int `json:"totalMinionsKilled"`

	// Position/Role
	Lane         string `json:"lane"`
	Role         string `json:"role"`
	TeamPosition string `json:"teamPosition"`

	// Placement (for game modes with rankings)
	Placement int `json:"placement"`
}

type GameStats struct {
	Participant      Participant
	GameDuration     uint32
	PerformanceScore float64
	Opponent         Participant
	Win              bool
}
