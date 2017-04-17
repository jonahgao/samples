//
// Created by jonah on 2017/4/14.
//

package sintervaltree

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type tendpoint int

func (te tendpoint) Compare(other Endpoint) int {
	return int(te) - int(other.(tendpoint))
}

type tinterval struct {
	left, right int
}

func (t tinterval) Left() Endpoint {
	return tendpoint(t.left)

}

func (t tinterval) Right() Endpoint {
	return tendpoint(t.right)
}

const tsize = 1000

func TestBasic(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// random inputs
	intervals := make([]Interval, 0, tsize)
	for i := 0; i < tsize; i++ {
		intervals = append(intervals, randInterval(rnd))
	}

	tree := NewTree(intervals)
	for i := 0; i < tsize; i++ {
		x := randEndPoint(rnd)
		res1 := tree.QueryPoint(tendpoint(x))
		res2 := queryPointNaive(intervals, tendpoint(x))
		if err := compareResults(res1, res2); err != nil {
			t.Logf("inputs: %v", intervals)
			t.Logf("query: %v", x)
			t.Logf("res1: %v", res1)
			t.Logf("res2: %v", res2)
			t.Fatal(err)
		}
	}
}

func BenchmarkIntervalTree(b *testing.B) {
	rnd := rand.New(rand.NewSource(1))

	// random inputs
	intervals := make([]Interval, 0, tsize)
	for i := 0; i < tsize; i++ {
		intervals = append(intervals, randInterval(rnd))
	}
	tree := NewTree(intervals)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := randEndPoint(rnd)
		res1 := tree.QueryPoint(tendpoint(x))
		_ = res1
	}
}

func BenchmarkNaive(b *testing.B) {
	rnd := rand.New(rand.NewSource(1))

	// random inputs
	intervals := make([]Interval, 0, tsize)
	for i := 0; i < tsize; i++ {
		intervals = append(intervals, randInterval(rnd))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		x := randEndPoint(rnd)
		res := queryPointNaive(intervals, tendpoint(x))
		_ = res
	}
}

func randEndPoint(rnd *rand.Rand) int {
	return rnd.Int() % 1000
}

func randInterval(rnd *rand.Rand) (result tinterval) {
	a, b := randEndPoint(rnd), randEndPoint(rnd)
	if a < b {
		result.left = a
		result.right = b
	} else {
		result.left = b
		result.right = a
	}
	return
}

func queryPointNaive(intervals []Interval, x Endpoint) (results []Interval) {
	for _, i := range intervals {
		if x.Compare(i.Left()) >= 0 && x.Compare(i.Right()) <= 0 {
			results = append(results, i)

		}
	}
	return
}

func queryNaive(intervals []Interval, q Interval) []Interval {
	//TODO:
	return nil
}

type resultSorter []Interval

func (rs resultSorter) Len() int {
	return len(rs)
}

func (rs resultSorter) Less(i, j int) bool {
	c := rs[i].Left().Compare(rs[j].Left())
	if c != 0 {
		return c < 0
	}
	return rs[i].Right().Compare(rs[j].Right()) < 0
}

func (rs resultSorter) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func compareResults(a []Interval, b []Interval) error {
	sort.Sort(resultSorter(a))
	sort.Sort(resultSorter(b))
	if len(a) != len(b) {
		return fmt.Errorf("length mismatch: %d != %d\n res1=%v\n res2=%v", len(a), len(b), a, b)
	}
	for i := 0; i < len(a); i++ {
		if a[i].Left().Compare(b[i].Left()) != 0 {
			return fmt.Errorf("left mismatch: %v != %v, at index %d\n res1=%v\n res2=%v", a[i].Left(), b[i].Left(), i, a, b)
		}
		if a[i].Right().Compare(b[i].Right()) != 0 {
			return fmt.Errorf("right mismatch: %v != %v, at index %d\n res1=%v\n res2=%v", a[i].Right(), b[i].Right(), i, a, b)
		}
	}
	return nil
}
