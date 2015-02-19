package core

import (
	"testing"
)

func PrepareSamples(sample1 *Sample, sample2 *Sample) {
	f1 := new(Feature)
	f1.FromString("1:1.0")
	f2 := new(Feature)
	f2.FromString("2:1.2")
	f3 := new(Feature)
	f3.FromString("5:2.2")
	f4 := new(Feature)
	f4.FromString("7:4.0")
	f5 := new(Feature)
	f5.FromString("11:12.12")
	f6 := new(Feature)
	f6.FromString("4:3.1")
	f7 := new(Feature)
	f7.FromString("72:20.2")

	sample1.AddFeature(*f1)
	sample1.AddFeature(*f2)
	sample1.AddFeature(*f3)
	sample1.AddFeature(*f4)
	sample1.AddFeature(*f5)

	sample2.AddFeature(*f1)
	sample2.AddFeature(*f2)
	sample2.AddFeature(*f6)
	sample2.AddFeature(*f7)
}

func TestNormalizer(t *testing.T) {
	sample1 := NewSample()
	sample2 := NewSample()
	PrepareSamples(sample1, sample2)

	samples := []*Sample{sample1, sample2}

	LinearNormalize(samples)
	if sample1.ToString(false) !=
		"0 1:0 2:0.010416666666666664 5:0.06250000000000001 7:0.15625 11:0.5791666666666666 " {
		t.Errorf("%s", sample1.ToString(false))
	}
	if sample2.ToString(false) !=
		"0 1:0 2:0.010416666666666664 4:0.10937500000000001 72:1 " {
		t.Errorf("%s", sample2.ToString(false))
	}
}

func TestLog10Normalizer(t *testing.T) {
	sample1 := NewSample()
	sample2 := NewSample()
	PrepareSamples(sample1, sample2)

	samples := []*Sample{sample1, sample2}

	Log10Normalize(samples)
	if sample1.ToString(false) !=
		"0 1:0 2:0.07918124604762482 5:0.3424226808222063 7:0.6020599913279624 11:1.0835026198302673 " {
		t.Errorf("%s", sample1.ToString(false))
	}
	if sample2.ToString(false) !=
		"0 1:0 2:0.07918124604762482 4:0.4913616938342727 72:1.3053513694466239 " {
		t.Errorf("%s", sample2.ToString(false))
	}
}

func TestAtanNormalizer(t *testing.T) {
	sample1 := NewSample()
	sample2 := NewSample()
	PrepareSamples(sample1, sample2)

	samples := []*Sample{sample1, sample2}

	AtanNormalize(samples)
	if sample1.ToString(false) !=
		"0 1:0.5 2:0.5577158767526089 5:0.7284005024398162 7:0.8440417392452613 11:0.9475923247149136 " {
		t.Errorf("%s", sample1.ToString(false))
	}
	if sample2.ToString(false) !=
		"0 1:0.5 2:0.5577158767526089 4:0.8013478156017629 72:0.9685098775965943 " {
		t.Errorf("%s", sample2.ToString(false))
	}
}
