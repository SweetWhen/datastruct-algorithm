package main

import (
	"fmt"	
	"math/rand"
	"time"
)

const (
	cutOff = 10
)

func main() {
	//generate a private rand
	r := rand.New(rand.NewSource(9))
	//set seed
	r.Seed(time.Now().UnixNano())
	
	//use r.Perm generate 0~n number
	a := r.Perm(100000)
	//use timestamp to statistics a process time-consuming incorrect maybe..
	start := time.Now()
	MyQsort(a, 0, len(a)-1)
	fmt.Printf("%d numbers Use %v \n", len(a), time.Since(start).String())

	var b = make([]int, 1000000)
	for  i := 0; i < len(b); i++ {
		b[i] = r.Int()
	}
	begin := time.Now()
	MyQsort(b, 0, len(b)-1)
	fmt.Printf("%d numbers Use %v \n", len(b), time.Since(begin).String())

}

func MyQsort(a []int, start, end int) {
	if start + cutOff < end {
		i := start
		j := end-1
		minValue := findSetMin(a, start, end)
		for {

			for i++; a[i] < minValue; i++  {
			}
			for j--; a[j] > minValue; j-- {
			}

			if i < j {
				a[i], a[j] = a[j], a[i]
			} else {
				break
			}
		}
		a[i], a[end-1] = a[end-1], a[i]
		MyQsort(a, start, i-1)
		MyQsort(a, i+1, end)
	} else {
		insertSort(a, start, end)
	}
} 

func findSetMin(a []int, start, end int) int {
	min := (start + end) / 2

	if a[start] > a[min] {
		a[start], a[min] = a[min], a[start]
	}
	if a[min] > a[end] {
		a[min], a[end] = a[end], a[min]
	}
	if a[start] > a[min] {
		a[start], a[min] = a[min], a[start]
	}
	
	//a[min] store in a[end-1]
	a[min], a[end-1] = a[end-1], a[min]
	
	return a[end-1]
		
}

func insertSort(a []int, start, end int) {
	if end - start > 0 {
		for i := start+1; i <= end; i++ {
			for j := i-1; j >= start; j-- {
				if a[j] > a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}
	}
}

func PrintArray(a []int, msg string) {
	fmt.Println(msg)
	for _, i := range a {
		fmt.Printf("%d ", i)
	}
	fmt.Println()	
}
