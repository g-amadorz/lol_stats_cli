package printer

import (
	"fmt"
	"lol_stats/internal/persistence"
)

type Cell struct {
	id    int
	score float64
}

type grid [][]Cell

func PrintPerformanceChart(performances []persistence.Performance) {

	cells := []Cell{}

	for _, p := range performances {
		// fmt.Printf("Lane: '%s', Score: %.1f\n", p.Participant.Lane, p.Score)

		cell := Cell{
			id:    p.Idx,
			score: p.Score,
		}
		cells = append(cells, cell)
	}

	grid := buildGrid(cells)

	for i := range len(grid) {
		for k := range len(grid[i]) {

			printCell(grid[i][k])
		}
		fmt.Println()
	}
}

// Fix Later but works for now not exact layout [0,1,2,3,4] etc
func buildGrid(cells []Cell) grid {
	g := make(grid, 4)
	for i := range g {
		g[i] = make([]Cell, 5)
	}
	v := 5
	for i := range 4 {
		for k := (v * i); k < v*(i+1); k++ {
			g[i][k%5] = Cell{
				cells[k].id,
				cells[k].score,
			}
		}
	}
	return g
}

func printCell(cell Cell) {
	val := cell.score
	id := cell.id

	var escape string
	switch {
	case val == 0:
		escape = "\033[48;2;100;100;100m\033[38;2;255;255;255m"
	case val > 0 && val < 5:
		escape = "\033[48;2;255;0;0m\033[38;2;0;0;0m"
	case val >= 5 && val < 10:
		escape = "\033[48;2;255;255;0m\033[38;2;0;0;0m"
	case val >= 10:
		escape = "\033[48;2;0;255;0m\033[38;2;0;0;0m"
	}

	var str string

	if id < 10 {
		str = fmt.Sprint("0", id)
	} else {
		str = fmt.Sprint(id)
	}

	fmt.Printf("%s  %s  \033[0m", escape, str)
}

func PrintParticipantStats(performance persistence.Performance) {

	p := performance.Participant

	resultColor := "\033[1;31m"
	result := "DEFEAT"
	if p.Win {
		resultColor = "\033[1;32m"
		result = "VICTORY"
	}

	fmt.Printf("\n%s════════════════════════════════════════════════════════\033[0m\n", resultColor)
	fmt.Printf("%s%s %s\033[0m\n", resultColor, result, resultColor)
	fmt.Printf("%s════════════════════════════════════════════════════════\033[0m\n\n", resultColor)

	fmt.Printf("Player: %s#%s\n", p.RiotIDGameName, p.RiotIDTagline)
	fmt.Printf("Champion: %s (Level %d)\n", p.ChampionName, p.ChampLevel)
	fmt.Printf("Position: %s\n\n", p.Lane)

	kda := "Perfect"
	if p.Deaths > 0 {
		kda = fmt.Sprintf("%.2f", float64(p.Kills+p.Assists)/float64(p.Deaths))
	}
	fmt.Printf("━━━ Combat Stats ━━━\n")
	fmt.Printf("  K/D/A:  %d / %d / %d  (KDA: %s)\n", p.Kills, p.Deaths, p.Assists, kda)
	fmt.Printf("  Damage Dealt:  %s\n", formatNumber(p.TotalDamageDealtToChampions))
	fmt.Printf("  Damage Taken:  %s\n", formatNumber(p.TotalDamageTaken))
	fmt.Printf("  CC Duration:   %.1fs\n\n", float64(p.TotalTimeCCDealt))

	// Economy
	fmt.Printf("━━━ Economy ━━━\n")
	fmt.Printf("  Gold Earned:  %s\n", formatNumber(p.GoldEarned))
	fmt.Printf("  CS:           %d\n", p.TotalMinionsKilled)
	csPerMin := 0.0
	if p.TimePlayed > 0 {
		csPerMin = float64(p.TotalMinionsKilled) / (float64(p.GameDuration) / 60.0)
	}

	fmt.Printf("  CS/min:       %.1f\n\n", csPerMin)

	// Vision & Support
	fmt.Printf("━━━ Vision & Support ━━━\n")
	fmt.Printf("  Vision Score:  %d\n", p.VisionScore)
	if p.TotalHealsOnTeammates > 0 {
		fmt.Printf("  Team Healing:  %s\n", formatNumber(p.TotalHealsOnTeammates))
	}

	fmt.Printf("\n")
}

func formatNumber(n int) string {
	str := fmt.Sprintf("%d", n)
	if len(str) <= 3 {
		return str
	}

	result := ""
	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(digit)
	}
	return result
}
