//
// Created by jonah on 2017/4/14.
//

package sintervaltree

import (
	"sort"
)

// // Intervals slice of Interval
// type Intervals []Interval

type node struct {
	mid   interface{} // middle point
	left  *node       // left child, intervals fully left of mid
	right *node       // right child, intervals fully right of mid

	overlapAsecLeft  []Interval // those intervals that overlap with mid endpoint, sort by their left endpoints ascending
	overlapDescRight []Interval // those intervals that overlap with mid endpoint, sort by their right endpoints descending
}

func newNode(mid interface{}, left, right *node, overlap []Interval) *node {
	n := &node{
		mid:   mid,
		left:  left,
		right: right,
	}

	if len(overlap) > 0 {
		n.overlapAsecLeft = make([]Interval, 0, len(overlap))
		n.overlapAsecLeft = append(n.overlapAsecLeft, overlap...)
		sort.Sort(leftAsecSorter(n.overlapAsecLeft))

		n.overlapDescRight = make([]Interval, 0, len(overlap))
		n.overlapDescRight = append(n.overlapAsecLeft, overlap...)
		sort.Sort(rightDescSorter(n.overlapDescRight))
	}

	return n
}

type leftAsecSorter []Interval

func (ls leftAsecSorter) Len() int {
	return len(ls)
}

func (ls leftAsecSorter) Less(i, j int) bool {
	return ls[i].Compare(ls[i].Left(), ls[j].Left()) < 0
}

func (ls leftAsecSorter) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

type rightDescSorter []Interval

func (rs rightDescSorter) Len() int {
	return len(rs)
}

func (rs rightDescSorter) Less(i, j int) bool {
	return rs[i].Compare(rs[i].Right(), rs[j].Right()) > 0
}

func (rs rightDescSorter) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
