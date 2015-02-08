package core

import (
	"github.com/golang/glog"

	"util"
)

func GetSimilarity(sampleA *Sample, sampleB *Sample) float64 {
	sampleA_feature_ids := util.NewHashSet()
	sampleB_feature_ids := util.NewHashSet()
	if len(sampleA.Features) == 0 || len(sampleB.Features) == 0 {
		glog.Errorf("Empty sample.\n")
		return 0.0
	}
	for _, f := range sampleA.Features {
		sampleA_feature_ids.Add(f.Id)
	}
	for _, f := range sampleB.Features {
		sampleB_feature_ids.Add(f.Id)
	}

	total_feature_num := util.Union(sampleA_feature_ids,
		sampleB_feature_ids).Len()
	equal_feature_num := 0.0
	for _, common_id := range util.Intersect(
		sampleA_feature_ids, sampleB_feature_ids).(util.HashSet).Elements() {
		if sampleA.Features[common_id.(int)].Value ==
			sampleB.Features[common_id.(int)].Value {
			equal_feature_num++
		}
	}
	return equal_feature_num / float64(total_feature_num)
}
