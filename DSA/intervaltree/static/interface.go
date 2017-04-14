//
// Created by jonah on 2017/4/14.
//

package sintervaltree

// Interval the interval
type Interval interface {
	// Left return interval's left endpoint
	Left() interface{}

	// Right return interval's right endpoint
	Right() interface{}

	// Compare compare two endpoint
	Compare(interface{}, interface{}) int
}

// Tree the interface of static interval tree
type Tree interface {
	// QueryPoint Find all intervals in the tree which contain the query point
	QueryPoint(interface{}) []Interval

	// Query Find all intervals in the tree which intersect the given interval
	Query(Interval) []Interval
}
