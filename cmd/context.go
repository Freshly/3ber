package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	contextCmd = &cobra.Command{
		Use:     "context",
		Aliases: []string{"context"},
		Short:   "Manage Kubernetes config contexts",
		Run: func(cmd *cobra.Command, args []string) {
			mustExist("kubectl")

			fmt.Printf("TODO\n")
			// kubectl config current-context
			// kubectl config get-contexts -o name
			// kubectl config set-context gke_freshly-staging_us-central1_staging-v2
		},
	}
)
