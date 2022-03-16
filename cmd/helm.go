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
	helmCmd = &cobra.Command{
		Use:   "helm",
		Short: "manage helm charts",
		Long:  `You can manually manage helm charts using these commands.`,
	}
	installCmdNamespaceFlag = ""
	installCmdSetFlags      = []string{}
	installCmd              = &cobra.Command{
		Use:     "install [RELEASE] [CHART]",
		Aliases: []string{"i"},
		Short:   "install a helm chart release",
		Long: `This command installs a chart archive.

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
	uninstallCmdNamespaceFlag = ""
	uninstallCmd              = &cobra.Command{
		Use:     "uninstall [RELEASE]",
		Aliases: []string{"u"},
		Short:   "uninstall a helm chart release",
		PreRun:  helmMustExist,
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
	installCmd.Flags().StringVarP(&installCmdNamespaceFlag, "namespace", "n", "", "install release into a target namespace")
	installCmd.Flags().StringArrayVarP(&installCmdSetFlags, "set", "s", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	listCmd.Flags().BoolVarP(&listCmdAllNamespacesFlag, "all-namespaces", "A", false, "list releases across all namespaces")
	listCmd.Flags().StringVarP(&listCmdNamespaceFlag, "namespace", "n", "", "list releases for a target namespace")
	uninstallCmd.Flags().StringVarP(&uninstallCmdNamespaceFlag, "namespace", "n", "", "uninstall release from target namespace")
	upgradeCmd.Flags().StringVarP(&upgradeCmdNamespaceFlag, "namespace", "n", "", "upgrade release in a target namespace")
	upgradeCmd.Flags().StringArrayVarP(&upgradeCmdSetFlags, "set", "s", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	helmCmd.AddCommand(installCmd)
	helmCmd.AddCommand(listCmd)
	helmCmd.AddCommand(uninstallCmd)
	helmCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(helmCmd)
}
