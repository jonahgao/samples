//
// Created by jonah on 2017/4/14.
//

package sintervaltree

type node struct {
	mid   EndPoint // middle point
	left  *node    // left child
	right *node    // right child

	overlapAsecLeft  Intervals // those intervals that overlap withi mid endpoint, sort by left endpoint ascending
	overlapDescRight Intervals // those intervals that overlap withi mid endpoint, sort by right endpoint descending
}

type tree struct {
	root *node
}

// NewTree construction new interval tree. complexity: o(nlogn)
func NewTree(intervals []Interval) Tree {
	return &tree{}
}

func (t *tree) QueryPoint(EndPoint) Intervals {
	//TODO:
	return nil
}

func (t *tree) Query(Interval) Intervals {
	//TODO:
	return nil
}
