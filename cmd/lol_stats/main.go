package main

import (
	"fmt"
	"lol_stats/internal/api"
)

func main() {
	id, _ := api.QueryAccount("FREE PALESTINE", "tox")

	fmt.Println(api.QueryMatches(id))
}
