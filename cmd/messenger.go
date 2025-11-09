package cmd

import (
	"doorayctl/result"
	"log/slog"
	"os"

	"github.com/dooray-go/dooray/openapi/messenger"
	"github.com/spf13/cobra"
)

var messengerCmd = &cobra.Command{
	Use:   "messenger [organizationMemberId] [message]",
	Short: "Send direct message",
	Long:  `Send a direct message to a Dooray organization member.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		log := slog.New(slog.NewTextHandler(os.Stdout, nil))

		env, err := GetEnv()
		if err != nil {
			log.Warn("Failed to get environment", "error", err)
			return
		}

		organizationMemberId := args[0]
		message := args[1]

		messengerClient := messenger.NewDefaultMessenger()
		sendResult, err := messengerClient.DirectSend(env.Token, &messenger.DirectSendRequest{
			Text:                 message,
			OrganizationMemberId: organizationMemberId,
		})
		if err != nil {
			log.Warn("Messenger Send Failed.", "error", err)
			return
		}

		err = result.PrintMessengerResult(sendResult)
		if err != nil {
			log.Warn("Report Failed.", "error", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(messengerCmd)
}