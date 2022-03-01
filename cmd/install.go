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
	installCmdNamespaceFlag = ""
	installCmdSetFlags      = []string{}
	installCmd              = &cobra.Command{
		Use:     "install [RELEASE_NAME] [CHART]",
		Aliases: []string{"i"},
		Short:   "install a helm chart release",
		Long: `This command installs a chart archive.

The RELEASE_NAME argument must be a unique release name.

The CHART argument must be a chart reference, a path to a packaged chart,
a path to an unpacked chart directory or a URL.`,
		PreRun: func(cmd *cobra.Command, args []string) {
			mustExist("helm")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				if err := cmd.Usage(); err != nil {
					fmt.Printf("usage error: %v\n", err)
				}
				os.Exit(1)
			}
			releaseName := args[0]
			chart := args[1]

			helmArgs := []string{"install", releaseName, chart}
			if installCmdNamespaceFlag != "" {
				helmArgs = append(helmArgs, "--create-namespace")
				helmArgs = append(helmArgs, "--namespace", installCmdNamespaceFlag)
			}
			for _, installCmdSetFlag := range installCmdSetFlags {
				helmArgs = append(helmArgs, "--set", installCmdSetFlag)
			}

			voice.Sayf("Installing Helm chart %s with release name %s", chart, releaseName)
			c := exec.Command("helm", helmArgs...)
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	installCmd.Flags().StringVarP(&installCmdNamespaceFlag, "namespace", "n", "", "install release into a target namespace")
	installCmd.Flags().StringArrayVarP(&installCmdSetFlags, "set", "s", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	rootCmd.AddCommand(installCmd)
}
