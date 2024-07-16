package run

import (
	"github.com/spf13/cobra"
)

var flags = struct {
	port int
}{}

var configFile = "etc/webservice.yaml"

func NewCommand() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run webService in console",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(configFile)
		},
	}
	runCmd.Flags().IntVarP(&flags.port, "port", "p", 8080, "listening port")

	return runCmd
}
