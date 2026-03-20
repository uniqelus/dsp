package signal

import (
	"math"
	"time"

	"github.com/uniqelus/dsp/internal/models"
)

type Generator struct {
	offset       float64
	startPhase   float64
	amplitude    float64
	angularFreq  float64
	samplingFreq float64
}

func NewGenerator(opts ...GeneratorOption) (*Generator, error) {
	options := defaultOptions()
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	return &Generator{
		offset:       options.offset,
		startPhase:   options.startPhase,
		amplitude:    options.amplitude,
		angularFreq:  2 * math.Pi * options.signalFreq,
		samplingFreq: options.samplingFreq,
	}, nil
}

func (g *Generator) Generate(duration time.Duration) []*models.Sample {
	totalSamples := int(math.Floor(duration.Seconds() * g.samplingFreq))
	samples := make([]*models.Sample, 0, totalSamples)

	for i := range totalSamples {
		sampleTime := float64(i) / g.samplingFreq

		samples = append(samples, g.generateSample(sampleTime))
	}

	return samples
}

func (g *Generator) generateSample(time float64) *models.Sample {
	value := g.amplitude*math.Cos(g.angularFreq*time+g.startPhase) + g.offset

	return &models.Sample{
		Value: value,
		Time:  time,
	}
}
