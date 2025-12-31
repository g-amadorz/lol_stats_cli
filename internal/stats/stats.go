package stats

import (
	"fmt"
	"lol_stats/internal/model"
)

type Cell struct {
	id    int
	score float64
}

type grid [][]Cell

func calculateScoreTop(model.Participant) float64 {

	return 10
}

func calculateScoreJungle(model.Participant) float64 {
	return 8
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

func calculateScore(player model.Participant) float64 {
	switch player.Lane {
	case "Top":
		return calculateScoreTop(player)
	case "Jungle":
		return calculateScoreJungle(player)

	case "Mid":
		return calculateScoreMid(player)

	case "Bot":
		return calculateScoreBot(player)

	case "Support":
		return calculateScoreSupport(player)

	default:
		return 0
	}

}

func PrintPerformanceChart(performances []model.Participant) {

	cells := []Cell{}

	for idx, performance := range performances {
		cell := Cell{
			id:    idx,
			score: calculateScore(performance),
		}
		cells = append(cells, cell)
	}

	grid := buildGrid(cells)

	for i := range len(grid) {
		for k := range len(grid[i]) {
			printCell(grid[i][k])
		}
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

	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 5:
		escape = "\033[1;30;47m"
	case val >= 5 && val < 10:
		escape = "\033[1;30;43m"
	case val >= 10:
		escape = "\033[1;30;42m"
	}

	str := "  %s "

	fmt.Printf(escape+str+"\033[0m", id)
}
