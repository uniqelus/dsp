package generate

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/uniqelus/dsp/internal/generators/signal"
	"github.com/uniqelus/dsp/internal/models"
)

var genSignalCmd = &cobra.Command{
	Use:   "signal",
	Short: "Generate harmonic discrete signal (cosine wave)",
	Long: `Generates an array of discrete signal samples

The output is saved to a CSV file with two columns: time (time in seconds) 
and value (signal value in Volts).`,
	Example: `
  dsp generate signal \
    --amplitude 5.0 \
    --offset 2.0 \
    --signal-frequency 100 \
    --sampling-frequency 10000 \
    --initial-phase 0.785398 \
    --simulation-time 100ms \
    --output 1_signal.csv`,
	RunE: func(cmd *cobra.Command, args []string) error {
		amplitude, _ := cmd.Flags().GetFloat64("amplitude")
		offset, _ := cmd.Flags().GetFloat64("offset")
		signalFreq, _ := cmd.Flags().GetFloat64("signal-frequency")
		samplingFreq, _ := cmd.Flags().GetFloat64("sampling-frequency")
		startPhase, _ := cmd.Flags().GetFloat64("initial-phase")
		simTime, _ := cmd.Flags().GetDuration("simulation-time")
		outputPath, _ := cmd.Flags().GetString("output")

		generator, err := signal.NewGenerator(
			signal.WithAmplitude(amplitude),
			signal.WithOffset(offset),
			signal.WithSignalFrequency(signalFreq),
			signal.WithSamplingFrequency(samplingFreq),
			signal.WithStartPhase(startPhase),
		)
		if err != nil {
			return err
		}

		samples := generator.Generate(simTime)

		if err := writeSamplesToCSV(samples, outputPath); err != nil {
			return fmt.Errorf("failed to write CSV: %w", err)
		}

		return nil
	},
}

func init() {
	genSignalCmd.Flags().
		Float64("amplitude", signal.DefaultAmplitude, "signal amplitude A [V]")
	genSignalCmd.Flags().
		Float64("offset", signal.DefaultOffset, "DC offset C [V]")
	genSignalCmd.Flags().
		Float64("signal-frequency", signal.DefaultSignalFreq, "harmonic signal frequency f [Hz]")
	genSignalCmd.Flags().
		Float64("sampling-frequency", signal.DefaultSamplingFreq, "sampling frequency Fs [Hz]")
	genSignalCmd.Flags().
		Float64("initial-phase", signal.DefaultStartPhase, "initial phase φ [rad]")
	genSignalCmd.Flags().
		Duration("simulation-time", 100*time.Millisecond, "simulation duration [ms/s]")
}

func writeSamplesToCSV(samples []*models.Sample, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"time", "value"}); err != nil {
		return err
	}

	for _, s := range samples {
		record := []string{
			strconv.FormatFloat(s.Time, 'g', -1, 64),
			strconv.FormatFloat(s.Value, 'g', -1, 64),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
