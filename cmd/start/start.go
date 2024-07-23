package start

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start webservice in daemon mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			runArgs := append([]string{"run"}, os.Args[2:]...)
			proc := exec.Command(os.Args[0], runArgs...)
			if err := proc.Start(); err != nil {
				fmt.Printf("fail to start webservice, error is: %v\n", err)
				return err
			}
			fmt.Printf("success to start webservice\n")
			return nil
		},
	}
	return startCmd
}
