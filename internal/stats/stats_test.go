package stats

import (
	"fmt"
	"testing"
)

func TestBuildGrid(t *testing.T) {
	tests := []struct {
		name  string
		cells []Cell
		want  grid
	}{
		{
			name: "builds 4x5 grid from 20 cells",
			cells: []Cell{
				{id: 0, score: 1.0}, {id: 1, score: 2.0}, {id: 2, score: 3.0}, {id: 3, score: 4.0}, {id: 4, score: 5.0},
				{id: 5, score: 6.0}, {id: 6, score: 7.0}, {id: 7, score: 8.0}, {id: 8, score: 9.0}, {id: 9, score: 10.0},
				{id: 10, score: 11.0}, {id: 11, score: 12.0}, {id: 12, score: 13.0}, {id: 13, score: 14.0}, {id: 14, score: 15.0},
				{id: 15, score: 16.0}, {id: 16, score: 17.0}, {id: 17, score: 18.0}, {id: 18, score: 19.0}, {id: 19, score: 20.0},
			},
			want: grid{
				{{id: 0, score: 1.0}, {id: 1, score: 2.0}, {id: 2, score: 3.0}, {id: 3, score: 4.0}, {id: 4, score: 5.0}},
				{{id: 5, score: 6.0}, {id: 6, score: 7.0}, {id: 7, score: 8.0}, {id: 8, score: 9.0}, {id: 9, score: 10.0}},
				{{id: 10, score: 11.0}, {id: 11, score: 12.0}, {id: 12, score: 13.0}, {id: 13, score: 14.0}, {id: 14, score: 15.0}},
				{{id: 15, score: 16.0}, {id: 16, score: 17.0}, {id: 17, score: 18.0}, {id: 18, score: 19.0}, {id: 19, score: 20.0}},
			},
		},
		{
			name: "builds grid with all zeros",
			cells: []Cell{
				{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0},
				{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0},
				{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0},
				{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0},
			},
			want: grid{
				{{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}},
				{{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}},
				{{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}},
				{{id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}, {id: 0, score: 0}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildGrid(tt.cells)

			// Check dimensions
			if len(got) != 4 {
				t.Errorf("Expected 4 rows, got %d", len(got))
			}

			for i, row := range got {
				if len(row) != 5 {
					t.Errorf("Row %d: expected 5 columns, got %d", i, len(row))
				}
			}

			// Check each cell
			for i := range got {
				for j := range got[i] {
					if got[i][j] != tt.want[i][j] {
						t.Errorf("Cell [%d][%d]: got %+v, want %+v", i, j, got[i][j], tt.want[i][j])
					}
				}
			}
		})
	}
}

func TestColors(t *testing.T) {
	fmt.Println("\033[1;30;42m GREEN \033[0m")
	// fmt.Println("\033[1;30;43m YELLOW \03           3[0m")
	fmt.Println("\033[1;30;47m WHITE \033[0m")
}
