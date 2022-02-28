package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "3ber",
		Short: "Freshly infrastructure management tool",
	}
)

func init() {
	viper.AutomaticEnv()
	rootCmd.PersistentFlags().BoolP("voice", "v", false, "enable voice synthesizer")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "disable logging to stdout")
	_ = viper.BindPFlag("3BER_VOICE", rootCmd.PersistentFlags().Lookup("voice"))
	_ = viper.BindPFlag("3BER_QUIET", rootCmd.PersistentFlags().Lookup("quiet"))

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
