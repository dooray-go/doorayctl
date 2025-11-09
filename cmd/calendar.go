package cmd

import (
	"doorayctl/result"
	"log/slog"
	"os"

	"github.com/dooray-go/dooray/openapi/calendar"
	"github.com/spf13/cobra"
)

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Manage calendars",
	Long:  `Manage Dooray calendars.`,
}

var calendarListCmd = &cobra.Command{
	Use:   "list",
	Short: "List calendars",
	Long:  `List all calendars in Dooray.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := slog.New(slog.NewTextHandler(os.Stdout, nil))

		env, err := GetEnv()
		if err != nil {
			log.Warn("Failed to get environment", "error", err)
			return
		}

		calendarClient := calendar.NewDefaultCalendar()
		calendarsResponse, err := calendarClient.GetCalendars(env.Token)
		if err != nil {
			log.Warn("Get Calendars Failed.", "error", err)
			return
		}

		err = result.PrintCalendarsResult(calendarsResponse)
		if err != nil {
			log.Warn("Get Calendars Failed.", "error", err)
			return
		}
	},
}

func init() {
	calendarCmd.AddCommand(calendarListCmd)
	rootCmd.AddCommand(calendarCmd)
}