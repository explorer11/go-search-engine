package searchengine

import "sort"

func topNValues(mapParam map[int]int, n int) map[int]int {
	topN := map[int]int{}

	keys := make([]int, 0, len(mapParam))
	for key, _ := range mapParam {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return mapParam[keys[i]] > mapParam[keys[j]]
	})

	i := 0
	for _, key := range keys {
		i++
		if i <= n {
			topN[key] = mapParam[key]
		} else {
			break
		}
	}

	return topN
}
