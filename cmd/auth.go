package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	authCmd = &cobra.Command{
		Use:     "auth",
		Aliases: []string{"auth"},
		Short:   "Authenticate to Google Cloud and populate Kubernetes config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("TODO\n")
			// gcloud auth login
			// gcloud container clusters get-credentials internal-v2 --zone us-central1-f --project freshly-internal
			// gcloud container clusters get-credentials production-v2 --region us-central1 --project freshly-production
			// gcloud container clusters get-credentials staging-test --region us-east4 --project freshly-staging
			// gcloud container clusters get-credentials staging-v2 --region us-central1 --project freshly-staging
		},
	}
)
