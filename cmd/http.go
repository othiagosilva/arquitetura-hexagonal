package cmd

import (
	"fmt"
	server2 "github.com/codeedu/go-hexagonal/adapters/web/server"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebserver()
		server.Service = &productService
		fmt.Println("Webserver has been started")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}