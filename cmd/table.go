package cmd

import (
	"github.com/spf13/cobra"
	"release-handler/internal/ui"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Release table helps devs deploy faster",
	Long: `A simple CLI tool that helps developers
			get a synchronized view of jira tickets and Pull Requests`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.CreateReleaseTable()
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
