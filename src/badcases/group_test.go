package badcases

import (
	"testing"

	"core"
)

func TestGrouping(t *testing.T) {
	sample1 := core.NewSample()
	sample2 := core.NewSample()
	sample3 := core.NewSample()
	sample4 := core.NewSample()
	sample5 := core.NewSample()

	f1 := new(core.Feature)
	f1.FromString("1:1.0")
	f2 := new(core.Feature)
	f2.FromString("2:1.2")
	f3 := new(core.Feature)
	f3.FromString("5:2.2")
	f4 := new(core.Feature)
	f4.FromString("7:4.0")
	f5 := new(core.Feature)
	f5.FromString("11:12.12")
	f6 := new(core.Feature)
	f6.FromString("4:3.1")
	f7 := new(core.Feature)
	f7.FromString("21:0.2")
	f8 := new(core.Feature)
	f8.FromString("22:43.1")
	f9 := new(core.Feature)
	f9.FromString("44:5.0")
	f10 := new(core.Feature)
	f10.FromString("59:10.9")
	f11 := new(core.Feature)
	f11.FromString("72:20.2")

	sample1.AddFeature(*f1)
	sample1.AddFeature(*f2)
	sample1.AddFeature(*f3)
	sample1.AddFeature(*f4)
	sample1.AddFeature(*f5)
	sample1.AddFeature(*f8)
	sample1.AddFeature(*f10)

	sample2.AddFeature(*f1)
	sample2.AddFeature(*f2)
	sample2.AddFeature(*f6)
	sample2.AddFeature(*f7)
	sample2.AddFeature(*f8)

	sample3.AddFeature(*f1)
	sample3.AddFeature(*f3)
	sample3.AddFeature(*f5)
	sample3.AddFeature(*f7)
	sample3.AddFeature(*f9)
	sample3.AddFeature(*f11)

	sample4.AddFeature(*f2)
	sample4.AddFeature(*f4)
	sample4.AddFeature(*f6)
	sample4.AddFeature(*f8)
	sample4.AddFeature(*f10)

	sample5.AddFeature(*f1)
	sample5.AddFeature(*f2)
	sample5.AddFeature(*f3)
	sample5.AddFeature(*f5)
	sample5.AddFeature(*f7)
	sample5.AddFeature(*f9)
	sample5.AddFeature(*f11)

	samples := []*core.Sample{sample1, sample2, sample3, sample4, sample5}

	ret := TopCloseCasePairs(samples, 10.0, 3)
	if len(ret) != 3 {
		t.Errorf("Wrong result length: %d. Should be 3.", len(ret))
	}

	expected_1 := "2 <-> 3 : 0.000000"
	if ret[0].ToString() != expected_1 {
		t.Errorf("Wrong result element: %s. Should be %s", ret[0].ToString(), expected_1)
	}

	expected_2 := "3 <-> 4 : 0.090909"
	if ret[1].ToString() != expected_2 {
		t.Errorf("Wrong result element: %s. Should be %s", ret[1].ToString(), expected_2)
	}

	expected_3 := "1 <-> 2 : 0.222222"
	if ret[2].ToString() != expected_3 {
		t.Errorf("Wrong result element: %s. Should be %s", ret[2].ToString(), expected_3)
	}
}
