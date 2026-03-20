package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/uniqelus/dsp/cmd/dsp/generate"
)

var rootCmd = &cobra.Command{
	Use:   "dsp",
	Short: "Command line tool for digital signals and noises",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("config", "", "", "path to tool configuration")

	rootCmd.AddCommand(generate.GenerateCommand)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}
