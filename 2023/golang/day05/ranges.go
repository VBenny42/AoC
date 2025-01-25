package day05

import (
	"sort"
)

type rangeType struct {
	start int
	end   int
}

type ranges struct {
	ranges []rangeType
}

func (r *ranges) union(other ranges) {
	switch {
	case len(r.ranges) == 0 && len(other.ranges) == 0:
		return
	case len(r.ranges) == 0 && len(other.ranges) != 0:
		r.ranges = other.ranges
	case len(r.ranges) != 0 && len(other.ranges) == 0:
		return
	default:
		r.ranges = append(r.ranges, other.ranges...)

		sort.Slice(r.ranges, func(i, j int) bool {
			return r.ranges[i].start < r.ranges[j].start
		})

		var newRanges []rangeType
		currentRange := r.ranges[0]

		for i := 1; i < len(r.ranges); i++ {
			if currentRange.end < r.ranges[i].start {
				newRanges = append(newRanges, currentRange)
				currentRange = r.ranges[i]
			} else {
				currentRange.end = max(r.ranges[i].end, currentRange.end)
			}
		}
		newRanges = append(newRanges, currentRange)

		r.ranges = newRanges
	}
}

func (r *ranges) intersection(other ranges) ranges {
	var newRanges []rangeType

	for _, r1 := range r.ranges {
		for _, r2 := range other.ranges {
			if r1.start < r2.end && r1.end > r2.start {
				newRanges = append(newRanges, rangeType{
					max(r1.start, r2.start),
					min(r1.end, r2.end),
				})
			}
		}
	}

	return ranges{newRanges}
}

func (r *ranges) difference(other ranges) ranges {
	newRanges := make([]rangeType, len(r.ranges))
	copy(newRanges, r.ranges)

	for _, otherRange := range other.ranges {
		var tempRanges []rangeType
		for _, r := range newRanges {
			if r.start >= otherRange.end || r.end <= otherRange.start {
				tempRanges = append(tempRanges, rangeType{r.start, r.end})
			} else {
				if r.start < otherRange.start {
					tempRanges = append(tempRanges, rangeType{r.start, otherRange.start})
				}
				if r.end > otherRange.end {
					tempRanges = append(tempRanges, rangeType{otherRange.end, r.end})
				}
			}
		}
		newRanges = tempRanges
	}

	return ranges{newRanges}
}
