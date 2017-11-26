package main

import (
	"os"

	"github.com/arschles/jar/pkg/commands/generate"
	"github.com/spf13/cobra"
)

// RootCmd returns the root command for gostatic
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jar",
		Short: "The simple static site generator.",
	}
	cmd.AddCommand(generate.Root())
	return cmd
}

func main() {
	cmd := newRootCmd()
	if err := cmd.Execute(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}
}
