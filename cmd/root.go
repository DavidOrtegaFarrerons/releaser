package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"release-handler/config"
)

var rootCmd = &cobra.Command{
	Use:   "release-handler",
	Short: "A CLI tool to help developers make easier releases",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	configFile := filepath.Join(homeDir, ".release-handler.yaml")

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("Config file not found, creating default config at:", configFile)

		viper.SetDefault(config.JiraApiKey, "")
		viper.SetDefault(config.JiraDomain, "")
		viper.SetDefault(config.JiraJQL, "")
		viper.SetDefault(config.JiraEmail, "")
		viper.SetDefault(config.AzureUserId, "")
		viper.SetDefault(config.AzureApiKey, "")
		viper.SetDefault(config.AzureOrganization, "")
		viper.SetDefault(config.AzureProject, "")
		viper.SetDefault(config.AzureRepositoryId, "")
		viper.SetDefault(config.TicketPrefix, "")
		viper.SetDefault(config.CorsDomain, "")

		if err := viper.SafeWriteConfigAs(configFile); err != nil {
			fmt.Println("Error writing default config:", err)
		}
	}

	viper.AutomaticEnv()
	if err := viper.BindEnv(config.JiraApiKey); err != nil {
		fmt.Println("Error binding JIRA_KEY:", err)
	}
	if err := viper.BindEnv(config.AzureApiKey); err != nil {
		fmt.Println("Error binding AZURE_KEY:", err)
	}
	if err := viper.BindEnv(config.TicketPrefix); err != nil {
		fmt.Println("Error binding TICKET_PREFIX:", err)
	}

	// Read config file
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Error reading config file:", err)
	}
}
