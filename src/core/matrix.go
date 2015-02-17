package core

import (
	"sort"
	"strconv"
	"strings"
	"util"
)

type Matrix struct {
	Data map[int64]*Vector
}

func NewMatrix() *Matrix {
	m := Matrix{}
	m.Data = make(map[int64]*Vector)
	return &m
}

func (m *Matrix) AddValue(k1, k2 int64, v float64) {
	_, ok := m.Data[k1]
	if !ok {
		m.Data[k1] = NewVector()
	}
	m.Data[k1].AddValue(k2, v)
}

func (m *Matrix) SetValue(k1, k2 int64, v float64) {
	row, ok := m.Data[k1]
	if !ok {
		row = NewVector()
		m.Data[k1] = row
	}
	row.SetValue(k2, v)
}

func (m *Matrix) GetValue(k1, k2 int64) float64 {
	row := m.GetRow(k1)
	if row == nil {
		return 0.0
	} else {
		return row.GetValue(k2)
	}
}

func (m *Matrix) GetRow(k1 int64) *Vector {
	row, ok := m.Data[k1]
	if !ok {
		return nil
	} else {
		return row
	}
}

func (m *Matrix) Scale(scale float64) *Matrix {
	ret := NewMatrix()
	for id, vi := range m.Data {
		ret.Data[id] = vi.Scale(scale)
	}
	return ret
}

func (m *Matrix) MultiplyVector(v *Vector) *Vector {
	// This is intended for l-by-m * m-by-1
	// For m-by-1 * 1-by-n, use OuterProduct in vector.go
	// Probably should just have a MatrixMultiply for everything
	ret := NewVector()
	for id, vi := range m.Data {
		ret.SetValue(id, v.Dot(vi))
	}
	return ret
}

func (m *Matrix) Trans() *Matrix {
	ret := NewMatrix()
	for rid, vi := range m.Data {
		for cid, w := range vi.Data {
			ret.SetValue(cid, rid, w)
		}
	}
	return ret
}

func (m *Matrix) ElemWiseAddMatrix(n *Matrix) *Matrix {
	ret := NewMatrix()
	for key, mi := range m.Data {
		ret.Data[key] = mi
	}
	for key, ni := range n.Data {
		if ret.GetRow(key) == nil {
			ret.Data[key] = ni
		} else {
			ret.Data[key] = ni.ElemWiseAddVector(ret.GetRow(key))
		}
	}
	return ret
}

func (m *Matrix) ToBytes() []byte {
	sb := util.StringBuilder{}

	keys := make([]int, len(m.Data))
	i := 0
	for k := range m.Data {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)

	for key := range keys {
		sb.Int(key)
		sb.Write(">")
		sb.WriteBytes(m.Data[int64(key)].ToBytes())
		sb.Write("\n")
	}
	return sb.Bytes()
}

func (m *Matrix) ToString() string {
	sb := util.StringBuilder{}

	keys := make([]int, len(m.Data))
	i := 0
	for k := range m.Data {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)

	for _, key := range keys {
		sb.Int(key)
		sb.Write(">")
		sb.Write(m.Data[int64(key)].ToString())
		sb.Write("\n")
	}
	return sb.String()
}

func (m *Matrix) FromString(buf string) {
	lines := strings.Split(buf, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		tks := strings.Split(line, ">")
		i, _ := strconv.Atoi(tks[0])
		row := NewVector()
		row.FromString(tks[1])
		m.Data[int64(i)] = row
	}
}
