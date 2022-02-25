package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "3ber",
		Short: "Freshly infrastructure management tool",
	}
)

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose logging and voice synthesizer")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "disable info logging and voice synthesizer")

	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(contextCmd)
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
