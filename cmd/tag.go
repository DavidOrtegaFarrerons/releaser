package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"release-handler/internal/release"
)

var environment string

var tagCmd = &cobra.Command{
	Use:   "tag [environment]",
	Short: "Creates a release tag",
	Long:  `A simple CLI command that helps a developer generate a release tag`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envInput := args[0]

		envMap := map[string]string{
			"rls":        "release",
			"release":    "release",
			"prd":        "prod",
			"prod":       "prod",
			"production": "prod",
		}

		normalizedEnv, exists := envMap[envInput]

		if !exists {
			fmt.Println("The env: " + envInput + " does not exist")
			os.Exit(1)
		}

		tag := release.GenerateTag(normalizedEnv)

		fmt.Println("Your new tag for the env " + normalizedEnv + " is:")
		fmt.Println(tag)
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.Flags().StringVarP(&environment, "env", "v", "", "Set for which env the tag should be generated")
}
