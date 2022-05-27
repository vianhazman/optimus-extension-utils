package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:          "utils",
		SilenceUsage: true,
	}
	rootCmd.AddCommand(windowCheckerCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
