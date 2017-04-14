//
// Created by jonah on 2017/4/14.
//

package sintervaltree

import "sort"

type tree struct {
	root *node
	cmp  Interval
}

// NewTree construction new interval tree. complexity: o(nlogn)
func NewTree(intervals []Interval) Tree {
	t := &tree{}
	if len(intervals) == 0 {
		return t
	}
	t.cmp = intervals[0]
	t.root = t.build(intervals)
	return t
}

func (t *tree) QueryPoint(x interface{}) []Interval {
	if t == nil || t.root == nil {
		return nil
	}
	// TODO:
	return nil
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

type endpointSorter struct {
	cmp    Interval
	inputs []interface{}
}

func (es endpointSorter) Len() int {
	return len(es.inputs)
}

func (es endpointSorter) Less(i, j int) bool {
	return es.cmp.Compare(es.inputs[i], es.inputs[j]) < 0
}

func (es endpointSorter) Swap(i, j int) {
	es.inputs[i], es.inputs[j] = es.inputs[j], es.inputs[i]
}

func (t *tree) partition(inputs []Interval) (mid interface{}, left, right, overlap []Interval) {
	// TODO: remove this check when stable
	if len(inputs) == 0 {
		panic("StaticIntervalTree::partition(): empty inputs")
	}

	endpoints := make([]interface{}, 0, len(inputs)*2)
	for _, i := range inputs {
		endpoints = append(endpoints, i.Left())
		endpoints = append(endpoints, i.Right())
	}
	sort.Sort(endpointSorter{cmp: t.cmp, inputs: endpoints})

	mid = endpoints[len(endpoints)/2]

	for _, i := range inputs {
		if t.cmp.Compare(i.Left(), mid) > 0 { // fully right
			left = append(left, i)
		} else if t.cmp.Compare(i.Right(), mid) < 0 { // fully left
			right = append(right, i)
		} else { // overlap with mid
			overlap = append(overlap, i)
		}
	}

	return
}
