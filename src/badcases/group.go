package badcases

import (
	"fmt"
	"sort"
	"sync"

	"core"
	"util"
)

type SamplePair struct {
	I        int
	J        int
	Distance float64
}

func (s SamplePair) ToString() string {
	return fmt.Sprintf("%d <-> %d : %f", s.I, s.J, s.Distance)
}

type byDistance []SamplePair

func (v byDistance) Len() int           { return len(v) }
func (v byDistance) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v byDistance) Less(i, j int) bool { return v[i].Distance < v[j].Distance }

func (v byDistance) ToString() string {
	sb := util.StringBuilder{}
	for _, s := range v {
		sb.Write(s.ToString())
		sb.Write("\n")
	}
	return sb.String()
}

// Returns Top close sample pairs.
// Each sample pair's distance should be less than threshold.
// Return at most N sample pairs.
// Return array is sorted by distance in ascending order.
func TopCloseCasePairs(samples []*core.Sample, threshold float64, N int) (
	ret []SamplePair) {
	L := len(samples)

	var wg sync.WaitGroup
	wg.Add(L * (L - 1) / 2)

	result := make(chan *SamplePair, L*(L-1)/2)
	count := 0
	for i, _ := range samples {
		for j := i + 1; j < L; j++ {
			go func(samples []*core.Sample, i int, j int,
				wg *sync.WaitGroup) {
				defer wg.Done()
				if i == j || i >= len(samples) ||
					j >= len(samples) {
					return
				}
				dist := core.GetSampleSimilarity(
					samples[i], samples[j])
				if dist <= threshold {
					result <- &SamplePair{
						I:        i,
						J:        j,
						Distance: dist,
					}
					count++
				}
			}(samples, i, j, &wg)
		}
	}
	wg.Wait()

	r := byDistance{}
	for i := 0; i < count; i++ {
		r = append(r, *(<-result))
	}
	close(result)

	sort.Sort(byDistance(r))

	if len(r) > N {
		return r[0:N]
	}

	return r
}
