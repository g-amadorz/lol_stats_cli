package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "lol_stats",
	Short:   "A brief description of your app",
	Long:    `A longer description...`,
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
