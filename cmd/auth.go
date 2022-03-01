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
	authCmdForceFlag = false
	authCmd          = &cobra.Command{
		Use:     "auth",
		Aliases: []string{"auth"},
		Short:   "Authenticate to Google Cloud and populate Kubernetes config",
		Run: func(cmd *cobra.Command, args []string) {
			mustExist("gcloud")
			mustExist("kubectl")

			if authCmdForceFlag {
				voice.Say("Forcefully reauthenticating.")
			} else if hasExpectedContexts() {
				voice.Say("I detected existing Kubernetes cluster credentials.")
				voice.Say("If you want to re-authenticate, use the -f or --force flag.")
				os.Exit(0)
			}

			voice.Say("Freshly hosts its infrastructure on Google Cloud. I will first help you authenticate.")
			go func() {
				voice.Say("Please navigate to your web browser and then login to your Freshly Okta account.")
			}()

			c := exec.Command("gcloud", "auth", "login")
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}
			voice.Say("Authentication was successful. I will now get the Kubernetes cluster credentials.")

			// TODO dynamically enumerate projects
			// TODO dynamically enumerate Kubernetes cluster credentials in each project
			c = exec.Command("gcloud", "container", "clusters", "get-credentials", "internal-v2",
				"--zone", "us-central1-f",
				"--project", "freshly-internal")
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}

			c = exec.Command("gcloud", "container", "clusters", "get-credentials", "production-v2",
				"--region", "us-central1",
				"--project", "freshly-production")
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}

			c = exec.Command("gcloud", "container", "clusters", "get-credentials", "staging-v2",
				"--region", "us-central1",
				"--project", "freshly-staging")
			if err := common.RunCommand(c, true); err != nil {
				fmt.Printf("command failed: %v\n", err)
				os.Exit(1)
			}

			voice.Say("Kubernetes cluster credentials were successfully retrieved.")
			voice.Say("You can inspect the credentials at $HOME/.kube/config")
		},
	}
)

func init() {
	authCmd.Flags().BoolVarP(&authCmdForceFlag, "force", "f", false, "re-authenticate even if credentials already exist")
}
