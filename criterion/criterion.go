package criterion

import (
	"math"
)

// type Metric int
//
// const (
// 	GINI Metric = iota
// 	ENTROPY
// )
//
// func (m Metric) String() string {
// 	switch m {
// 	case GINI:
// 		return "Gini"
// 	case ENTROPY:
// 		return "Entropy"
// 	default:
// 		return "Gini"
// 	}
// }

func xlog2(x float64) float64 {
	return math.Log(x) / math.Log(2)
}

func UniqueCounts(targets []interface{}) map[string]int {
	results := make(map[string]int)
	for _, target := range targets {
		if _, ok := results[target.(string)]; ok {
			results[target.(string)]++
		} else {
			results[target.(string)] = 1
		}
	}
	return results
}

type Metric func([]interface{}) float64

func GiniImpurity(targets []interface{}) float64 {
	nData := float64(len(targets))
	counts := UniqueCounts(targets)
	imp := 0.
	for k1 := range counts {
		c1 := float64(counts[k1]) / nData
		for k2 := range counts {
			if k1 == k2 {
				continue
			}
			c2 := float64(counts[k2]) / nData
			imp += c1 * c2
		}
	}
	return imp
}

func Entropy(targets []interface{}) float64 {
	nData := float64(len(targets))
	counts := UniqueCounts(targets)
	ent := 0.
	for k := range counts {
		p := float64(counts[k]) / nData
		ent = ent - p*xlog2(p)
	}
	return ent
}
