package core

import (
	"testing"
)

// Same samples return similarity as 1.0
func TestSimilarity1(t *testing.T) {
	sample1 := NewSample()
	sample2 := NewSample()

	f1 := new(Feature)
	f1.FromString("1:1.0")
	f2 := new(Feature)
	f2.FromString("2:1.2")
	f3 := new(Feature)
	f3.FromString("5:1.2")

	sample1.AddFeature(*f1)
	sample1.AddFeature(*f2)
	sample1.AddFeature(*f3)
	sample2.AddFeature(*f1)
	sample2.AddFeature(*f2)
	sample2.AddFeature(*f3)

	if GetSimilarity(sample1, sample2) != 1.0 {
		t.Error("Similarity wrong.")
	}
}

// Empty sample return similarity as 0.0
func TestSimilarity2(t *testing.T) {
	sample1 := NewSample()
	sample2 := NewSample()

	f1 := new(Feature)
	f1.FromString("1:1.0")
	sample1.AddFeature(*f1)

	if GetSimilarity(sample1, sample2) != 0.0 {
		t.Error("Similarity wrong.")
	}
}

func TestSimilarity3(t *testing.T) {
	sample1 := NewSample()
	sample2 := NewSample()

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

	if GetSimilarity(sample1, sample2) != 0.2857142857142857 {
		t.Errorf("Similarity wrong.")
	}
}
