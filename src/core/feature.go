package core

import (
	"strconv"
	"strings"
	"util"
)

type FeatureType int

var FeatureTypeEnum = struct {
	DISCRETE_FEATURE   FeatureType
	CONTINUOUS_FEATURE FeatureType
}{0, 1}

func GetFeatureType(key string) FeatureType {
	if key[0] == '#' {
		return FeatureTypeEnum.DISCRETE_FEATURE
	} else {
		return FeatureTypeEnum.CONTINUOUS_FEATURE
	}
}

type Feature struct {
	Id    int64
	Value float64
}

func (f *Feature) ToString() string {
	sb := util.StringBuilder{}
	sb.Int64(f.Id)
	sb.Write(":")
	sb.Float(f.Value)
	sb.Write(" ")
	return sb.String()
}

func (f *Feature) FromString(buf string) {
	kv := strings.Split(strings.Trim(buf, " "), ":")
	key, _ := strconv.ParseInt(kv[0], 10, 64)
	value, _ := strconv.ParseFloat(kv[1], 64)
	f.Id = key
	f.Value = value
}
