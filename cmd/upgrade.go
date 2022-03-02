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
	upgradeCmdNamespaceFlag = ""
	upgradeCmdSetFlags      = []string{}
	upgradeCmd              = &cobra.Command{
		Use:     "upgrade [RELEASE] [CHART]",
		Aliases: []string{"up"},
		Short:   "upgrade a helm chart release",
		Long: `This command upgrades a chart archive.

The RELEASE argument must be a unique release name.

The CHART argument must be a chart reference, a path to a packaged chart,
a path to an unpacked chart directory or a URL.`,
		PreRun: helmMustExist,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				if err := cmd.Usage(); err != nil {
					fmt.Printf("usage error: %v\n", err)
				}
				os.Exit(1)
			}
			releaseName := args[0]
			chart := args[1]

			helmArgs := []string{"upgrade", releaseName, chart}
			if upgradeCmdNamespaceFlag != "" {
				helmArgs = append(helmArgs, "--create-namespace")
				helmArgs = append(helmArgs, "--namespace", upgradeCmdNamespaceFlag)
			}
			for _, upgradeCmdSetFlag := range upgradeCmdSetFlags {
				helmArgs = append(helmArgs, "--set", upgradeCmdSetFlag)
			}

			voice.Sayf("Upgrading Helm chart %s with release name %s", chart, releaseName)
			c := exec.Command("helm", helmArgs...)
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	upgradeCmd.Flags().StringVarP(&upgradeCmdNamespaceFlag, "namespace", "n", "", "upgrade release in a target namespace")
	upgradeCmd.Flags().StringArrayVarP(&upgradeCmdSetFlags, "set", "s", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	rootCmd.AddCommand(upgradeCmd)
}
