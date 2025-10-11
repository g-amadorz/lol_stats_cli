package stats

import (
	"fmt"
	"lol_stats/internal/model"
)

type Cell struct {
	id    string
	score float64
}

type grid [][]Cell

func PrintPerformanceChart(games []model.GameStats) {
	grid := buildGrid(games)

	for i := range len(grid) {
		for k := range len(grid[i]) {
			printCell(grid[i][k])
		}
	}
}

// Fix Later but works for now not exact layout [0,1,2,3,4] etc
func buildGrid(games []model.GameStats) grid {
	g := make(grid, 4)
	for i := range g {
		g[i] = make([]Cell, 5)
	}
	v := 5
	for i := range 4 {
		for k := (v * i); k < v*(i+1); k++ {
			g[i][k%5] = Cell{
				games[k].ID,
				games[k].PerformanceScore,
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

// Printing individual stats
func PrintGame(id string) {
	fmt.Println("Print individual game")
}
