//
// Created by jonah on 2017/4/14.
//

package sintervaltree

// Endpoint endpoint of interval
type Endpoint interface {
	Compare(Endpoint) int
}

// Interval the interval
type Interval interface {
	// Left return interval's left endpoint
	Left() Endpoint

	// Right return interval's right endpoint
	Right() Endpoint
}

// Tree the interface of static interval tree
type Tree interface {
	// QueryPoint Find all intervals in the tree which contain the query point
	QueryPoint(Endpoint) []Interval

	// Query Find all intervals in the tree which intersect the given interval
	Query(Interval) []Interval
}
