/*
source code for heap, for fun only, we don't use it, but container/heap
Init -> Push/Pop -> Fix -> Remove
*/

package myheap

import (	
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{}) //add x as element Len()
	Pop() interface{}  //remove and return element Len() -1
}

//init a slice or array to become heap
//complexity is O(n) where n = h.Len()

func Init(h Interface) {
	n := h.Len()
	for i := n/2 -1; i >= 0; i-- { //update from the bottom to top
		down(h, i, n)
	}
}

func down(h Interface, i_, n int) bool {
	i := i_
	for {
		jLeft := 2*i + 1
		if jLeft >= n || jLeft < 0 {
			break
		}
		j := jLeft //j contain the smallest
		if jRight := jLeft+1; jRight < n && h.Less(jRight, jLeft) {
			j = jRight //update j
		}
		if !h.Less(j, i) { //i is the one we need to find a position for it
			break //the smallest one isn't less than i, so i don't need to move any more
		}
		h.Swap(i, j)

		i = j //after i move to j, we need i contain the insert one		
	}

	return i > i_
}

//remove elem in the heap
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n >= 0 && n != i {
		h.Swap(i, n)
		if !down(h,i,n) {
			up(h, i)
		}
	}
	return h.Pop()
}


//Push the element x to the heap, complexity is O(log(n))
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

//up a elem to association position in a heap
func up(h Interface, j int) {
	for {
		i := (j-1)/2 //parent
		if i == j || h.Less(i, j) { 
			break //parent is less the me, so break
		}
		h.Swap(i,j)
		j = i //now i up to parent position
	}
}

//move the minimum element to the last heap
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	if n >= 0 {
		h.Swap(0, n)
		down(h, 0, n)
		return h.Pop()
	} else { 
		// a nil heap
		return h
	}
	
}

//re-establishes the heap after the element at indes has changed
//complexity is O(log(n))
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

 
