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
		fmt.Printf("Lane: '%s', Score: %.1f\n", p.Participant.Lane, p.Score)

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

	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 5:
		escape = "\033[1;30;47m"
	case val >= 5 && val < 10:
		escape = "\033[1;30;43m"
	case val >= 10:
		escape = "\033[1;30;42m"
	}

	str := "  %d "

	fmt.Printf(escape+str+"\033[0m", id)
}
