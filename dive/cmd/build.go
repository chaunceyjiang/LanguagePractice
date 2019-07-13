package cmd

import (
	"dive/runtime"
	"dive/utils"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:                "build [any valid `docker build` arguments]",
	Short:              "Builds and analyzes a docker image from a Dockerfile (this is a thin wrapper for the `docker build` command).",
	DisableFlagParsing: true,
	Run:                doBuildCmd,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

// doBuildCmd implements the steps taken for the build command
func doBuildCmd(cmd *cobra.Command, args []string) {
	defer utils.Cleanup()

	initLogging()

	runtime.Run(runtime.Options{
		BuildArgs:  args,
		ExportFile: exportFile,
	})
}
