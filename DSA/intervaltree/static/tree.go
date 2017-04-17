//
// Created by jonah on 2017/4/14.
//

package sintervaltree

import "sort"

// endpoint comparer
type comparer interface {
	Compare(interface{}, interface{}) int
}

type tree struct {
	root *node
}

// NewTree construction new interval tree. complexity: o(nlogn)
func NewTree(intervals []Interval) Tree {
	t := &tree{}
	if len(intervals) == 0 {
		return t
	}
	t.root = t.build(intervals)
	return t
}

func (t *tree) QueryPoint(x Endpoint) []Interval {
	if t == nil || t.root == nil {
		return nil
	}
	return t.root.queryPoint(x)
}

func (t *tree) Query(q Interval) []Interval {
	if t == nil || t.root == nil {
		return nil
	}
	// TODO:
	return nil
}

func (t *tree) build(intervals []Interval) *node {
	mid, left, right, overlap := t.partition(intervals)

	var leftNode, rightNode *node
	if len(left) > 0 {
		leftNode = t.build(left)
	}
	if len(right) > 0 {
		rightNode = t.build(right)
	}

	return newNode(mid, leftNode, rightNode, overlap)
}

type endpointSorter []Endpoint

func (es endpointSorter) Len() int {
	return len(es)
}

func (es endpointSorter) Less(i, j int) bool {
	return es[i].Compare(es[j]) < 0
}

func (es endpointSorter) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (t *tree) partition(inputs []Interval) (mid Endpoint, left, right, overlap []Interval) {
	// TODO: remove this check when stable
	if len(inputs) == 0 {
		panic("StaticIntervalTree::partition(): empty inputs")
	}

	endpoints := make([]Endpoint, 0, len(inputs)*2)
	for _, i := range inputs {
		endpoints = append(endpoints, i.Left())
		endpoints = append(endpoints, i.Right())
	}
	sort.Sort(endpointSorter(endpoints))

	mid = endpoints[len(endpoints)/2]

	for _, i := range inputs {
		if i.Left().Compare(mid) > 0 { // fully right
			right = append(right, i)
		} else if i.Right().Compare(mid) < 0 { // fully left
			left = append(left, i)
		} else { // overlap with mid
			overlap = append(overlap, i)
		}
	}

	return
}
