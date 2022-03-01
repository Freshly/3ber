package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/freshly/3ber/pkg/voice"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "3ber",
		Short: "Freshly infrastructure management tool",
	}
	mustExistInstallMap = map[string]string{
		"gcloud":  "https://cloud.google.com/sdk/docs/install",
		"helm":    "https://helm.sh/docs/intro/install/",
		"kubectl": "https://kubernetes.io/docs/tasks/tools/",
	}
	openBrowserCommandMap = map[string]string{
		"windows": "explorer",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
)

func init() {
	viper.AutomaticEnv()
	rootCmd.PersistentFlags().BoolP("voice", "v", false, "enable voice synthesizer")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "disable info logging")
	_ = viper.BindPFlag("VOICE", rootCmd.PersistentFlags().Lookup("voice"))
	_ = viper.BindPFlag("QUIET", rootCmd.PersistentFlags().Lookup("quiet"))
}

func mustExist(command string) string {
	path, err := exec.LookPath(command)
	if err != nil {
		voice.Say(fmt.Sprintf("Command '%s' not found, please install it.", command))

		if installURL, ok := mustExistInstallMap[command]; ok {
			if browserCommand, ok2 := openBrowserCommandMap[runtime.GOOS]; ok2 {
				_, err := exec.Command(browserCommand, installURL).CombinedOutput()
				if err != nil {
					voice.Say(fmt.Sprintf("Please navigate to %s", installURL))
				} else {
					voice.Say("I am navigating your default browser to the installation instructions.")
				}
			}
		}
		os.Exit(1)
	}
	return path
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
