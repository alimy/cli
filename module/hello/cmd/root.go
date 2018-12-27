package cmd

import (
	"github.com/alimy/cli/cmd"
	"github.com/alimy/cli/module/hello"
	"github.com/spf13/cobra"
)

func init() {
	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "sample sub-command for cli",
		Run:   helloRun,
	}

	// Register helloCmd as sub-command
	cmd.Register(helloCmd)
}

func helloRun(cmd *cobra.Command, args []string) {
	hello.Start()
}
