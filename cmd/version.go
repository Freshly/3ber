package cmd

import (
	"fmt"

	"github.com/freshly/3ber/pkg/version"
	"github.com/freshly/3ber/pkg/voice"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver"},
		Short:   "print the program version",
		Run: func(cmd *cobra.Command, args []string) {
			message := fmt.Sprintf("version %s, git commit %s", version.Version, version.GitCommit)
			if viper.GetBool("QUIET") {
				fmt.Println(message)
			} else {
				voice.Say(message)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
