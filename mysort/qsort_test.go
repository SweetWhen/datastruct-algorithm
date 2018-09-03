package main

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

var a []int
var b []int

func init() {
	r := rand.New(rand.NewSource(10))
	fmt.Println("init func print...")
	r.Seed(time.Now().UnixNano())
	a = r.Perm(100000)
	b = r.Perm(10000000)
}

func TestQsortFunction(t *testing.T) {
	a := []int {3, 5, 6, 2, 1, 33, 55, 10, 12, 78, 32, 24, 60}
	b := []int {1, 2, 3, 5, 6, 10, 12, 24, 32, 33, 55, 60, 78}
	MyQsort(a, 0, len(a)-1)
	fmt.Println("sort result:", a)
	if ok := sliceEqua(a, b); !ok {
		t.Errorf("sort func failed")
	}
}

func TestQsortPerfomance1(t *testing.T) {
	t.Logf("0~100000 numbers qsor..\n")
	MyQsort(a, 0, len(a)-1)
}

func TestQsortPerfomance2(t *testing.T) {
	t.Logf("0~1000000 numbers qsor..\n")
	MyQsort(b, 0, len(b)-1)
}

func sliceEqua(dst, src []int) bool {
	if len(dst) != len(src) {
		return false
	}
	
	for i, val := range dst {
		if val != src[i] {
			return false
		}
	}

	return true
}
