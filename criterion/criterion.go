package criterion

import "math"

func xlog2(x float64) float64 {
	return math.Log(x) / math.Log(2)
}

func uniqueCounts(targets []string) map[string]int {
	var results map[string]int
	for _, target := range targets {
		if _, ok := results[target]; ok {
			results[target]++
		} else {
			results[target] = 1
		}
	}
	return results
}

func GiniImpurity(targets []string) float64 {
	nData := float64(len(targets))
	counts := uniqueCounts(targets)
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

func Entropy(targets []string) float64 {
	nData := float64(len(targets))
	counts := uniqueCounts(targets)
	ent := 0.
	for k := range counts {
		p := float64(counts[k]) / nData
		ent = ent - p*xlog2(p)
	}
	return ent
}
