package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"release-handler/internal/httpserver"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs a server with a FE that displays release information",
	Long: `This commands runs a server that will host a page where you can see
			all the information as in the cli, with live updates and more functionality`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting server...")
		if err := httpserver.Start(); err != nil {
			log.Fatalf("Failed to start server %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
