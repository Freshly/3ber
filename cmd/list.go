package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/freshly/3ber/pkg/common"
	"github.com/freshly/3ber/pkg/voice"
	"github.com/spf13/cobra"
)

var (
	listCmdAllNamespacesFlag = false
	listCmdNamespaceFlag     = ""
	listCmd                  = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list helm chart releases",
		Long:    "This command lists all of the releases for a specified namespace (uses current namespace context if namespace not specified).",
		PreRun:  helmMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			helmArgs := []string{"list"}

			if listCmdAllNamespacesFlag {
				voice.Say("Listing Helm chart releases in all namespaces")
				helmArgs = append(helmArgs, "--all-namespaces")
			} else if listCmdNamespaceFlag != "" {
				voice.Sayf("Listing Helm chart releases in namespace %s", listCmdNamespaceFlag)
				helmArgs = append(helmArgs, "--namespace", listCmdNamespaceFlag)
			} else {
				voice.Say("Listing Helm chart releases for current namespace context")
			}

			c := exec.Command("helm", helmArgs...)
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	listCmd.Flags().BoolVarP(&listCmdAllNamespacesFlag, "all-namespaces", "A", false, "list releases across all namespaces")
	listCmd.Flags().StringVarP(&listCmdNamespaceFlag, "namespace", "n", "", "list releases for a target namespace")
	rootCmd.AddCommand(listCmd)
}
