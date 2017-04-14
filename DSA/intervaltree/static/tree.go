//
// Created by jonah on 2017/4/14.
//

package sintervaltree

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
	// TODO:
	return nil
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

	return
}
