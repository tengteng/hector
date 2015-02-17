package core

import (
	// "fmt"
	"math"
	"testing"
)

func TestMatrix(t *testing.T) {
	a := NewMatrix()
	precision := 1e-9

	a.AddValue(3, 4, 1.78)

	if math.Abs(a.GetValue(3, 4)-1.78) > precision {
		t.Error("Get wrong value after set value")
	}

	a.AddValue(3, 4, -1.1)

	if math.Abs(a.GetValue(3, 4)-0.68) > precision {
		t.Error("Add value wrong")
	}

	b := NewMatrix()

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.SetValue(int64(i), int64(j), 1.0)
		}
	}

	c := b.Scale(2.0)

	if math.Abs(c.GetValue(7, 8)-2.0) > precision {
		t.Error("scale function error")
	}
}

func TestMatrixSaveLoad(t *testing.T) {
	m := NewMatrix()

	m.AddValue(0, 0, 1.23)
	m.AddValue(1, 6, 2.03)
	m.AddValue(2, 5, 3.14159265358979323846264338)
	m.AddValue(2, 4, 4.0)
	m.AddValue(3, 3, 5.20)
	m.AddValue(4, 2, 6.98)
	m.AddValue(6, 1, 7.212312343674)
	m.AddValue(99, 101, 8.88)

	expected := "0>0:1.23|\n1>6:2.03|\n2>5:3.141592653589793|4:4|\n3>3:5.2|\n4>2:6.98|\n6>1:7.212312343674|\n99>101:8.88|\n"
	str := m.ToString()
	if str != expected {
		t.Errorf("Wrong matrix output string: expected:\n%s\n but got:\n%s\n", expected, str)
	}

	mm := NewMatrix()
	mm.FromString(str)

	if mm.GetValue(0, 0) != 1.23 {
		t.Errorf("Wrong matrix parsed from string: expected\n1.23\n but got:\n%f\n", mm.GetValue(0, 0))
	}
	if mm.GetValue(1, 6) != 2.03 {
		t.Errorf("Wrong matrix parsed from string: expected\n2.03\n but got:\n%f\n", mm.GetValue(1, 6))
	}
	if mm.GetValue(2, 5) != 3.14159265358979323846264338 {
		t.Errorf("Wrong matrix parsed from string: expected\n3.14159265358979323846264338\n but got:\n%f\n", mm.GetValue(2, 5))
	}
	if mm.GetValue(2, 4) != 4.0 {
		t.Errorf("Wrong matrix parsed from string: expected\n4.0\n but got:\n%f\n", mm.GetValue(2, 4))
	}
	if mm.GetValue(3, 3) != 5.20 {
		t.Errorf("Wrong matrix parsed from string: expected\n5.20\n but got:\n%f\n", mm.GetValue(3, 3))
	}
	if mm.GetValue(4, 2) != 6.98 {
		t.Errorf("Wrong matrix parsed from string: expected\n6.98\n but got:\n%f\n", mm.GetValue(4, 2))
	}
	if mm.GetValue(6, 1) != 7.212312343674 {
		t.Errorf("Wrong matrix parsed from string: expected\n7.212312343674\n but got:\n%f\n", mm.GetValue(6, 1))
	}
	if mm.GetValue(99, 101) != 8.88 {
		t.Errorf("Wrong matrix parsed from string: expected\n8.88\n but got:\n%f\n", mm.GetValue(99, 101))
	}
}
