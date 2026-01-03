package cmd

import (
	"log"
	"lol_stats/internal/api"
	"lol_stats/internal/persistence"
	"lol_stats/internal/printer"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const gamename = "FREE PALESTINE"
const tag = "tox"

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		load, err := cmd.Flags().GetBool("load")
		if err != nil {
			log.Fatal("Error parsing `load` flag")
		}

		if err != nil {
			log.Fatal(err)
		}

		if load {
			apiKey := LoadApiKey()

			account, err := api.QueryAccount(gamename, tag, apiKey)

			if err != nil {
				log.Fatal(err)
			}
			performances := persistence.QueryPerformances(account, apiKey)

			persistence.SaveGames(performances)
		}

		performances, err := persistence.LoadGames()

		if err != nil {
			log.Fatal("Error parsing `game` flag")
		}

		game, err := cmd.Flags().GetInt("game")

		if err != nil {
			log.Fatal(err)
		}

		if game > 0 {
			printer.PrintParticipantStats(performances[game])
		} else {
			printer.PrintPerformanceChart(performances)
		}

	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().BoolP("load", "l", false, "Show past 20 games")
	statsCmd.Flags().IntP("game", "g", 0, "Show indexed game")
}

func LoadApiKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load env file")

	}

	apiKey := os.Getenv("API_KEY")

	return apiKey
}
