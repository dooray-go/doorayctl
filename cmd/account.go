package cmd

import (
	"doorayctl/result"
	"log/slog"
	"os"

	"github.com/dooray-go/dooray/openapi/account"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account [organizationId]",
	Short: "Get account members",
	Long:  `Retrieve members from a specific Dooray organization.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := slog.New(slog.NewTextHandler(os.Stdout, nil))

		env, err := GetEnv()
		if err != nil {
			log.Warn("Failed to get environment", "error", err)
			return
		}

		organizationId := args[0]

		accountClient := account.NewDefaultAccount()
		members, err := accountClient.GetMembers(env.Token, organizationId, "")
		if err != nil {
			log.Warn("GetMembers Failed.", "error", err)
			return
		}

		err = result.PrintAccountResult(members)
		if err != nil {
			log.Warn("Report Failed.", "error", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
