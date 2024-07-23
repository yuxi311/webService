package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yuxi311/webService/cmd/run"
	"github.com/yuxi311/webService/cmd/start"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "webservice",
		Short: "webservice entrypoint",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.AddCommand(run.NewCommand())
	rootCmd.AddCommand(start.NewCommand())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Cmd Execute Error: %v", err)
	}
}

func main() {
	Execute()
	fmt.Println("This is first")
}
