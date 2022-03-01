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
	uninstallCmdNamespaceFlag = ""
	uninstallCmd              = &cobra.Command{
		Use:     "uninstall [RELEASE_NAME]",
		Aliases: []string{"u"},
		Short:   "uninstall a helm chart release",
		PreRun: func(cmd *cobra.Command, args []string) {
			mustExist("helm")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				if err := cmd.Usage(); err != nil {
					fmt.Printf("usage error: %v\n", err)
				}
				os.Exit(1)
			}
			releaseName := args[0]

			helmArgs := []string{"uninstall", releaseName}
			if uninstallCmdNamespaceFlag != "" {
				helmArgs = append(helmArgs, "--namespace", uninstallCmdNamespaceFlag)
			}

			voice.Sayf("Uninstalling Helm chart with release name %s", releaseName)
			c := exec.Command("helm", helmArgs...)
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	uninstallCmd.Flags().StringVarP(&uninstallCmdNamespaceFlag, "namespace", "n", "", "uninstall release from target namespace")
	rootCmd.AddCommand(uninstallCmd)
}
