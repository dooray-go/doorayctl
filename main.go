package main

import (
	"doorayctl/result"
	"encoding/json"
	"errors"
	"github.com/dooray-go/dooray/openapi/account"
	"github.com/dooray-go/dooray/openapi/calendar"
	"github.com/dooray-go/dooray/openapi/messenger"
	"log/slog"
	"os"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if len(os.Args) < 3 {
		return
	}
	command, err := validateAndGetSubCommand(os.Args[1])
	if err != nil {
		log.Warn("invalid command", "command", os.Args[1])
		return
	}

	env, err := getEnv()
	if err != nil {
		return
	}

	switch command {
	case "account":
		if len(os.Args) > 2 {
			accountClient := account.NewDefaultAccount()
			members, err := accountClient.GetMembers(env.Token, os.Args[2], "")
			if err != nil {
				log.Warn("GetMembers Failed.", "error", err)
				return
			}

			err = result.PrintAccountResult(members)
			if err != nil {
				log.Warn("Report Failed.", "error", err)
				return
			}
		}
	case "messenger":
		if len(os.Args) > 2 {
			messengerClient := messenger.NewDefaultMessenger()
			sendResult, err := messengerClient.DirectSend(env.Token, &messenger.DirectSendRequest{
				Text:                 os.Args[3],
				OrganizationMemberId: os.Args[2],
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
		}

	case "calendar":
		if len(os.Args) > 2 {
			if os.Args[2] == "list" {
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
			} else {
				log.Warn("invalid command", "command", os.Args[2])
			}
		}
	}

}

var subCommands = map[string]struct{}{
	"account":   {},
	"messenger": {},
	"project":   {},
	"calendar":  {},
}

func validateAndGetSubCommand(sub string) (string, error) {

	if _, ok := subCommands[sub]; ok {
		return sub, nil
	}
	return "", errors.New("no sub command" + sub)
}

type DoorayEnv struct {
	Token string `json:"token"`
}

func getEnv() (*DoorayEnv, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(dir + "/.dooray/config")
	if err != nil {
		return nil, err
	}

	var doorayEnv DoorayEnv
	err = json.Unmarshal(file, &doorayEnv)
	if err != nil {
		return nil, errors.New(".dooray/config file parsing error")
	}

	return &doorayEnv, nil
}
