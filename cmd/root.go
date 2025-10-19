package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "lol_stats",
	Short:   "CLI tool to view League of Legends stats made in Go!",
	Long:    `CLI tool made using Go alongside Cobra in order to keep track of League of Legends stats`,
	Version: "0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from lol_stats!")
	},
}

func Execute(apiKey string) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
