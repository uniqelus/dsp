package signal

import "fmt"

const (
	DefaultAmplitude    float64 = 1.0
	DefaultOffset       float64 = 0.0
	DefaultSignalFreq   float64 = 50.0
	DefaultStartPhase   float64 = 0.0
	DefaultSamplingFreq float64 = 10000.0
)

type options struct {
	amplitude    float64
	offset       float64
	signalFreq   float64
	samplingFreq float64
	startPhase   float64
}

func defaultOptions() *options {
	return &options{
		amplitude:    DefaultAmplitude,
		offset:       DefaultOffset,
		signalFreq:   DefaultSignalFreq,
		startPhase:   DefaultStartPhase,
		samplingFreq: DefaultSamplingFreq,
	}
}

type GeneratorOption func(*options) error

func WithAmplitude(v float64) GeneratorOption {
	return func(o *options) error {
		if v < 0 {
			return fmt.Errorf("amplitude cannot be negative: %f", v)
		}
		o.amplitude = v
		return nil
	}
}

func WithOffset(v float64) GeneratorOption {
	return func(o *options) error {
		o.offset = v
		return nil
	}
}

func WithSignalFrequency(v float64) GeneratorOption {
	return func(o *options) error {
		if v <= 0 {
			return fmt.Errorf("signal frequency must be positive: %f Hz", v)
		}
		o.signalFreq = v
		return nil
	}
}

func WithStartPhase(v float64) GeneratorOption {
	return func(o *options) error {
		o.startPhase = v
		return nil
	}
}

func WithSamplingFrequency(v float64) GeneratorOption {
	return func(o *options) error {
		if v <= 0 {
			return fmt.Errorf("sampling frequency must be positive: %f Hz", v)
		}
		o.samplingFreq = v
		return nil
	}
}
