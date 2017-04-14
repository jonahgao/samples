//
// Created by jonah on 2017/4/14.
//

package sintervaltree

// EndPoint interval's endpoint
type EndPoint interface {
	// Compare compare to other endpoint
	Compare(EndPoint) int
}

// Interval the interval
type Interval interface {
	// Left return interval's left endpoint
	Left() EndPoint

	// Right return interval's right endpoint
	Right() EndPoint
}

// Intervals slice of Interval
type Intervals []Interval

// Tree the interface of static interval tree
type Tree interface {
	// QueryPoint Find all intervals in the tree which contain the query point
	QueryPoint(EndPoint) Intervals

	// Query Find all intervals in the tree which intersect the given interval
	Query(Interval) Intervals
}
