package core

import (
	"fmt"

	"github.com/golang/glog"

	"util"
)

func GetSampleSimilarity(sampleA *Sample, sampleB *Sample) float64 {
	if len(sampleA.Features) == 0 || len(sampleB.Features) == 0 {
		glog.Errorf("Empty sample.\n")
		return 0.0
	}

	sampleA_mapbased := sampleA.ToMapBasedSample()
	sampleB_mapbased := sampleB.ToMapBasedSample()
	sampleA_feature_ids := util.NewHashSet()
	sampleB_feature_ids := util.NewHashSet()

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
		sampleA_feature_ids, sampleB_feature_ids).(*util.HashSet).Elements() {
		fmt.Println(common_id)
		if sampleA_mapbased.Features[common_id.(int64)] ==
			sampleB_mapbased.Features[common_id.(int64)] {
			equal_feature_num++
		}
	}
	return equal_feature_num / float64(total_feature_num)
}

// func GetDatasetSimilarity(datasetA *DataSet, datasetB *DataSet, strict bool,
// 	threshold float64) float64 {
// 	if !strict {
// 		RemoveLowFreqFeatures(datasetA, threshold)
// 		RemoveLowFreqFeatures(datasetB, threshold)
// 	}
// }
