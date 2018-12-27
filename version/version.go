package version

import (
	"fmt"
	"github.com/alimy/cli/cmd"
	"github.com/spf13/cobra"
)

var (
	Version    = "0.0.0"
	APIVersion = "v1alphal"

	// BuildGitHash Value will be set during build
	BuildGitHash = "Not provided"

	// BuildTime Value will be set during build
	BuildTime = "Not provided"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "version of application",
		Long:  "version infomation for application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s (APIVersion:%s)\nBuildTime:%s\nBuildGitSHA:%s\n",
				Version, APIVersion, BuildTime, BuildGitHash)
		},
	}

	// Register versionCmd as sub-command
	cmd.Register(versionCmd)
}
