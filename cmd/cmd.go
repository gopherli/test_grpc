package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"test_grpc/client"
	"test_grpc/server"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "start and stop",
	Long:  "input cmd help for grpc cmd list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start server or client",
	Long:  `start grpc server or client`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			start(args[0])
		} else {
			fmt.Println("lack of args")
		}
	},
}

func start(name string) {
	switch name {
	case "serve":
		server.StartServe()
	case "client":
		client.StartClient()
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
