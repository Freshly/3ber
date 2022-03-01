package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/freshly/3ber/pkg/common"
	"github.com/spf13/cobra"
)

var (
	// TODO make context detection more dynamic, if possible
	expectedContexts = []string{
		"gke_freshly-internal_us-central1-f_internal-v2",
		"gke_freshly-production_us-central1_production-v2",
		"gke_freshly-staging_us-central1_staging-v2",
	}
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

func hasExpectedContexts() bool {
	contexts := getContextsOrDie()
	missingContexts := []string{}
	for _, expectedContext := range expectedContexts {
		hasContext := false
		for _, context := range contexts {
			if expectedContext == context {
				hasContext = true
				break
			}
		}
		if !hasContext {
			missingContexts = append(missingContexts, expectedContext)
		}
	}
	if len(missingContexts) > 0 {
		for _, missingContext := range missingContexts {
			fmt.Printf("Detected missing Kubernetes context: %s\n", missingContext)
		}
		return false
	}
	return true
}

func getContextsOrDie() []string {
	c := exec.Command("kubectl", "config", "get-contexts", "-o", "name")
	contexts, err := common.RunCommandOutputArray(c)
	if err != nil {
		fmt.Printf("command failed: %v\n", err)
		os.Exit(1)
	}
	return contexts
}
