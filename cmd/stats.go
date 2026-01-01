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
			log.Fatal("Error parsing `all` flag")
		}

		// history, err := persistence.GetConfigPath("history.json")

		if err != nil {
			log.Fatal(err)
		}

		if load {
			apiKey := "RGAPI-d93ba2b3-d94b-4738-9d81-6e27cf28d6d8"

			account, err := api.QueryAccount(gamename, tag, apiKey)

			if err != nil {
				log.Fatal(err)
			}
			performances := persistence.QueryPerformances(account, apiKey)

			persistence.SaveGames(performances)
		}

		performances, err := persistence.LoadGames()

		if err != nil {
			log.Fatal(err)
		}

		printer.PrintPerformanceChart(performances)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().BoolP("load", "l", false, "Show past 20 games")

}

func LoadApiKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load env file")

	}

	apiKey := os.Getenv("API_KEY")

	log.Println(apiKey)

	return apiKey
}
