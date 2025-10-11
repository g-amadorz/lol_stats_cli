package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		showAll, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal("Error parsing `all` flag")
		}

		game, _ := cmd.Flags().GetString("game")
		if showAll {
			fmt.Println("Where I am going to show last 20 games")
		} else if game != "" {
			fmt.Println("Printing the individual games stats")
		} else {
			fmt.Println("Show description")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().BoolP("all", "a", false, "Show past 20 games")
	serveCmd.Flags().String("game", "", "Specify a game")

}
