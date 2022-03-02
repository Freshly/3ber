package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/freshly/3ber/pkg/common"
	"github.com/freshly/3ber/pkg/voice"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultArgoNamespace    = "argocd"
	defaultArgoOutputFormat = "yaml"
)

var (
	argoCmd = &cobra.Command{
		Use:   "argo",
		Short: "manage the argo continuous delivery pipeline",
		Long: `Argo automates deploying and upgrading Helm charts whenever one or more
Kubernetes manifest specifications change in a git repository.

Argo can also be configured to populate Helm value override flags with build-
related events. For example, a Cloud Build might overwrite a Helm value
override triggering an automatic chart upgrade.`,
	}
	deleteArgoAppNamespaceFlag = ""
	deleteArgoAppCmd           = &cobra.Command{
		Use:     "delete [APP_NAME]",
		Aliases: []string{"del"},
		Short:   "delete argo application",
		PreRun:  kubectlMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				if err := cmd.Usage(); err != nil {
					fmt.Printf("usage error: %v\n", err)
				}
				os.Exit(1)
			}
			appName := args[0]

			voice.Say("Deleting the Argo application resource")
			kubeArgs := []string{"delete", fmt.Sprintf("application/%s", appName)}
			if deleteArgoAppNamespaceFlag != "" {
				kubeArgs = append(kubeArgs, "--namespace", deleteArgoAppNamespaceFlag)
			}

			c := exec.Command("kubectl", kubeArgs...)
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
	getArgoAppNamespaceFlag = ""
	getArgoAppOutputFlag    = ""
	getArgoAppCmd           = &cobra.Command{
		Use:    "get [APP_NAME]",
		Short:  "get argo application",
		PreRun: kubectlMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				if err := cmd.Usage(); err != nil {
					fmt.Printf("usage error: %v\n", err)
				}
				os.Exit(1)
			}
			appName := args[0]

			voice.Say("Fetching the Argo application resource")
			kubeArgs := []string{"get", fmt.Sprintf("application/%s", appName)}
			if getArgoAppNamespaceFlag != "" {
				kubeArgs = append(kubeArgs, "--namespace", getArgoAppNamespaceFlag)
			}
			if viper.GetBool("QUIET") {
				kubeArgs = append(kubeArgs, "--output", "jsonpath={}")
			} else {
				kubeArgs = append(kubeArgs, "--output", getArgoAppOutputFlag)
			}

			c := exec.Command("kubectl", kubeArgs...)
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
	listArgoAppsNamespaceFlag     = ""
	listArgoAppsAllNamespacesFlag = false
	listArgoAppsWatchFlag         = false
	listArgoAppsCmd               = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list argo applications",
		PreRun:  kubectlMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			voice.Say("Fetching the Argo application resources")

			kubeArgs := []string{"get", "application"}
			if listArgoAppsAllNamespacesFlag {
				kubeArgs = append(kubeArgs, "--all-namespaces")
			} else if listArgoAppsNamespaceFlag != "" {
				kubeArgs = append(kubeArgs, "--namespace", listArgoAppsNamespaceFlag)
			}
			if listArgoAppsWatchFlag {
				kubeArgs = append(kubeArgs, "--watch")
			}

			c := exec.Command("kubectl", kubeArgs...)
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	deleteArgoAppCmd.Flags().StringVarP(&deleteArgoAppNamespaceFlag, "namespace", "n", defaultArgoNamespace, "delete application from a target namespace")
	getArgoAppCmd.Flags().StringVarP(&getArgoAppNamespaceFlag, "namespace", "n", defaultArgoNamespace, "get application from a target namespace")
	getArgoAppCmd.Flags().StringVarP(&getArgoAppOutputFlag, "output", "o", defaultArgoOutputFormat, "output format, one of: name|json|jsonpath={}|yaml")
	listArgoAppsCmd.Flags().StringVarP(&listArgoAppsNamespaceFlag, "namespace", "n", defaultArgoNamespace, "list applications for a target namespace")
	listArgoAppsCmd.Flags().BoolVarP(&listArgoAppsAllNamespacesFlag, "all-namespaces", "A", false, "list applications across all namespaces")
	listArgoAppsCmd.Flags().BoolVarP(&listArgoAppsWatchFlag, "watch", "w", false, "after listing the applications, watch for changes")
	rootCmd.AddCommand(argoCmd)
	argoCmd.AddCommand(deleteArgoAppCmd)
	argoCmd.AddCommand(getArgoAppCmd)
	argoCmd.AddCommand(listArgoAppsCmd)
}
