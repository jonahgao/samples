//
// Created by jonah on 2017/4/14.
//

package sintervaltree

import "sort"

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
		n.overlapDescRight = append(n.overlapDescRight, overlap...)
		sort.Sort(rightDescSorter(n.overlapDescRight))
	}

	return n
}

func (n *node) queryOverlap(cmp comparer, x interface{}, left bool) (result []Interval) {
	if left {
		for _, i := range n.overlapAsecLeft {
			if cmp.Compare(i.Left(), x) > 0 {
				break
			} else {
				result = append(result, i)
			}
		}
	} else {
		for _, i := range n.overlapDescRight {
			if cmp.Compare(i.Right(), x) < 0 {
				break
			} else {
				result = append(result, i)
			}
		}
	}
	return
}

func (n *node) queryPoint(cmp comparer, x interface{}) (result []Interval) {
	c := cmp.Compare(x, n.mid)
	if c < 0 { // left
		if n.left != nil {
			result = append(result, n.left.queryPoint(cmp, x)...)
		}
		result = append(result, n.queryOverlap(cmp, x, true)...)
	} else if c > 0 { // right
		result = append(result, n.queryOverlap(cmp, x, false)...)
		if n.right != nil {
			result = append(result, n.right.queryPoint(cmp, x)...)
		}
	} else { // x is mid
		result = append(result, n.overlapAsecLeft...)
	}
	return
}

func (n *node) query(q Interval, cmp comparer) []Interval {
	// TODO:
	return nil
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
