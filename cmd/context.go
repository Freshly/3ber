package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/freshly/3ber/pkg/common"
	"github.com/freshly/3ber/pkg/voice"
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
		Aliases: []string{"ctx"},
		Short:   "manage kubernetes cluster contexts",
		Long: `Manage Kubernetes cluster contexts. Effective context management enables
effortless switching between different environments.`,
	}
	currentContextCmd = &cobra.Command{
		Use:     "current",
		Aliases: []string{"cur"},
		Short:   "Get current Kubernetes cluster context",
		PreRun:  kubectlMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			voice.Say("The current Kubernetes cluster context is defined in $HOME/.kube/config")
			c := exec.Command("kubectl", "config", "current-context")
			contexts, err := common.RunCommandOutputArray(c)
			if err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(contexts[0])
		},
	}
	listContextsCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all available Kubernetes cluster contexts",
		PreRun:  kubectlMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			voice.Say("The available Kubernetes cluster contexts are defined in $HOME/.kube/config")
			contexts := getContextsOrDie()
			for _, context := range contexts {
				fmt.Println(context)
			}
		},
	}
	setContextCmd = &cobra.Command{
		Use:    "set [CONTEXT]",
		Short:  "Set current Kubernetes cluster context",
		PreRun: kubectlMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				if err := cmd.Usage(); err != nil {
					fmt.Printf("usage error: %v\n", err)
				}
				os.Exit(1)
			}
			desiredContext := args[0]

			existingContexts := getContextsOrDie()
			filteredContexts := []string{}
			for _, existingContext := range existingContexts {
				if strings.Contains(existingContext, desiredContext) {
					filteredContexts = append(filteredContexts, existingContext)
				}
			}

			switch {
			case len(filteredContexts) > 1:
				voice.Say("I found more than one matching context, so I will abort:")
				for _, filteredContext := range filteredContexts {
					fmt.Println(filteredContext)
				}
				os.Exit(1)
			case len(filteredContexts) == 1:
				desiredContext = filteredContexts[0]
			case len(filteredContexts) == 0:
				voice.Say("I could not find the requested context, so I will abort.")
				os.Exit(1)
				// TODO add force option maybe
				// voice.Say("To create a new context, you can specify the -f or --force flag")
			}

			voice.Say("I will set the current Kubernetes cluster context in $HOME/.kube/config")
			c := exec.Command("kubectl", "config", "use-context", desiredContext)
			results, err := common.RunCommandOutputArray(c)
			if err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(results[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(contextCmd)
	contextCmd.AddCommand(currentContextCmd)
	contextCmd.AddCommand(listContextsCmd)
	contextCmd.AddCommand(setContextCmd)
}

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
