package internal

import (
	"sort"
)

type Pair struct {
	Start, End int
}

func Collect(data []Interval) []ResultInterval {

	collect := map[Pair][]string{}
	periods := []int{}
	for _, period := range data {
		periods = append(periods, period.Start, period.End)
	}
	sort.Slice(periods, func(i, j int) bool {
		return periods[i] < periods[j]
	})
	result := []ResultInterval{}
	pLen := len(periods)
	for end := 1; end < pLen; end++ {
		pStart := periods[end-1]
		pEnd := periods[end]
		for _, interval := range data {
			if pStart >= interval.Start && pStart < interval.End {
				current := Pair{pStart, pEnd}
				temp := []string{}
				if names, ok := collect[current]; ok {
					temp = names
				}
				temp = append(temp, interval.Name)
				collect[current] = temp
			}
		}
	}
	for key, value := range collect {
		result = append(result, ResultInterval{key.Start, key.End, value})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Start < result[j].Start
	})
	return result
}
