package generate

import "github.com/spf13/cobra"

var GenerateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate discrete signal and noise data",
	Long: `Provides tools for generating arrays of discrete signal 
samples and noise sequences with configurable parameters.

Generated data is saved to a CSV file containing pairs of (time; value).`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	GenerateCommand.PersistentFlags().
		StringP("output", "o", "signal.csv", "path to output CSV file for saving samples")
	GenerateCommand.AddCommand(genSignalCmd)
}
