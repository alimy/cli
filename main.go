package main

import (
	"github.com/alimy/cli/cmd"
	_ "github.com/alimy/cli/version"
)

func main() {
	// Setup root cli command of application
	cmd.Setup(
		"cli",                      // command name
		"a sample cli application", // command short describe
		"a sample cli application base on go module style", // command long describe
	)

	// Execute start application
	cmd.Execute()
}
