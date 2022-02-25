package cmd

import (
	"fmt"

	"github.com/freshly/3ber/pkg/version"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"version"},
		Short:   "Print program version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s (%s)\n", version.Version, version.GitCommit)
		},
	}
)
