package core

import (
	"math"
)

const FEATURE_INF = 99999999.9

// Linear Normalization.
// Formula: x -> (x - min) / (max - min)
func LinearNormalize(samples []*Sample) {
	min := math.Inf(+1)
	max := math.Inf(-1)
	for _, s := range samples {
		for _, f := range s.Features {
			if min > f.Value {
				min = f.Value
			} else if max < f.Value {
				max = f.Value
			}
		}
	}
	max_dist := max - min
	if max_dist <= 0 {
		return
	}
	for i := 0; i < len(samples); i++ {
		for j := 0; j < len(samples[i].Features); j++ {
			samples[i].Features[j].Value = (samples[i].Features[j].Value - min) / max_dist
		}
	}
}

// Log Normalization.
// Formula: x -> Log10(x)
func Log10Normalize(samples []*Sample) {
	for i := 0; i < len(samples); i++ {
		for j := 0; j < len(samples[i].Features); j++ {
			if samples[i].Features[j].Value > 0 {
				samples[i].Features[j].Value = math.Log10(samples[i].Features[j].Value)
			} else if samples[i].Features[j].Value < 0 {
				samples[i].Features[j].Value = -math.Log10(-samples[i].Features[j].Value)
			} else {
				samples[i].Features[j].Value = FEATURE_INF
			}
		}
	}
}

// Atan Normalization.
// Formula: x -> atan(x) * 2 / Pi
func AtanNormalize(samples []*Sample) {
	for i := 0; i < len(samples); i++ {
		for j := 0; j < len(samples[i].Features); j++ {
			samples[i].Features[j].Value = math.Atan(samples[i].Features[j].Value) * 2 / math.Pi
		}
	}
}
