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

type tinterval struct {
	left, right int
}

func (t tinterval) Left() interface{} {
	return t.left
}

func (t tinterval) Right() interface{} {
	return t.right
}

func (t tinterval) Compare(lh, rh interface{}) int {
	return lh.(int) - rh.(int)
}

const tsize = 100000

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
		res1 := tree.QueryPoint(x)
		res2 := queryPointNaive(intervals, x)
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
		res1 := tree.QueryPoint(x)
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
		res := queryPointNaive(intervals, x)
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

func queryPointNaive(intervals []Interval, x interface{}) (results []Interval) {
	for _, i := range intervals {
		if i.Compare(x, i.Left()) >= 0 && i.Compare(x, i.Right()) <= 0 {
			// fmt.Printf("add %v-%v\n", i.Left(), i.Right())
			results = append(results, i)
		} else {
			// fmt.Printf("skip %v-%v\n", i.Left(), i.Right())
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
	c := rs[i].Compare(rs[i].Left(), rs[j].Left())
	if c != 0 {
		return c < 0
	}
	return rs[i].Compare(rs[i].Right(), rs[j].Right()) < 0
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
		if a[i].Compare(a[i].Left(), b[i].Left()) != 0 {
			return fmt.Errorf("left mismatch: %v != %v, at index %d\n res1=%v\n res2=%v", a[i].Left(), b[i].Left(), i, a, b)
		}
		if a[i].Compare(a[i].Right(), b[i].Right()) != 0 {
			return fmt.Errorf("right mismatch: %v != %v, at index %d\n res1=%v\n res2=%v", a[i].Right(), b[i].Right(), i, a, b)
		}
	}
	return nil
}
