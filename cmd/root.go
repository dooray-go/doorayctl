package cmd

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "doorayctl",
	Short: "Dooray CLI tool",
	Long:  `A command line interface for interacting with Dooray services.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// DoorayEnv represents the Dooray environment configuration
type DoorayEnv struct {
	Token string `json:"token"`
}

// GetEnv retrieves the Dooray environment configuration from the config file
func GetEnv() (*DoorayEnv, error) {
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