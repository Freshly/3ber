package cmd

import (
	"fmt"

	"github.com/freshly/3ber/pkg/version"
	"github.com/freshly/3ber/pkg/voice"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"version"},
		Short:   "Print program version",
		Run: func(cmd *cobra.Command, args []string) {
			voice.Say(fmt.Sprintf("version %s, git commit %s", version.Version, version.GitCommit))
		},
	}
)
