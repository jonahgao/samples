//
// Created by jonah on 2017/4/14.
//

package sintervaltree

import (
	"math/rand"
	"testing"
	"time"
)

type tinterval struct {
	left, right int
}

func (t tinterval) Left() interface{} {
	return left
}

func (t tinterval) Right() interface{} {
	return right
}

func (t tinterval) Compare(lh, rh interface{}) int {
	return lh.(int) - rh.(int)
}

const tsize = 100

func randInterval(rnd *rand.Rand) (result tinterval) {
	a, b := rnd.Int(), rnd.Int()
	if a < b {
		result.left = a
		result.right = b
	} else {
		result.left = b
		result.right = a
	}
	return
}

func queryPointNaive(intervals Intervals, x interface{}) Intervals {
	//TODO:
	return nil
}

func queryNaive(intervals Intervals, q Interval) Intervals {
	//TODO:
	return nil
}

func TestBasic(t *testing.T) {
	intervals := make([]Interval, 0, tsize)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < tsize; i++ {
		intervals = append(intervals, randInterval(rnd))
	}
	tree := NewTree(intervals)
	_ = tree
	// TODO:
}
