package tsp

import (
	"math"
	"sync"
	"concurrency-9/data_types"
	// "fmt"
)

// Merge_Sort, a parallel version of mergesort which brings down running time to O(n).
//
//   Input: s i.e. data_types.Graph_edge
//   Output: s, sorted i.e. data_types.Graph_edge
func Merge_Sort(s []data_types.Graph_edge) []data_types.Graph_edge {
	
	// nearest power of 2 to len(s)
	var k float64 = math.Ceil(math.Log2(float64(len(s))))
	
	// number of numbers to be added to the slice
	var diff int = int(math.Exp2(k) - float64(len(s)))

	// looping to add the deficit
	for i:=0;i<diff;i++ {
		s = append(s, data_types.Graph_edge{-1,-1,-1})
	}

	// call normal mergesort for smaller arrays
	if(len(s) < 1024) {
		normal_mergesort(s)
	} else {
		parallel_mergesort(s)
	}

	return s[diff:]
	// fmt.Printf("%v", s)
}

func normal_mergesort(s []data_types.Graph_edge) {

	// straight-forward implementation of mergesort

	if len(s) > 1 {
		middle := len(s) / 2
		// Split in Middle and continue
		normal_mergesort(s[:middle])
		normal_mergesort(s[middle:])
		merge(s, middle)
	}
}


func parallel_mergesort(s []data_types.Graph_edge) {
	// fmt.Println("using parallel")

	if len(s) > 1 {
		var len int = len(s)
		var middle int = len / 2

		var wg sync.WaitGroup
		wg.Add(2)

		// parallely merging the smaller slices
		// defers to make-sure completion
		go func() {
			defer wg.Done()
			parallel_mergesort(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallel_mergesort(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

func merge(s []data_types.Graph_edge, middle int) {
	helper := make([]data_types.Graph_edge, len(s))
	copy(helper, s)

	var tempLeft int = 0
	var tempRight int = middle
	var current int = 0
	high := len(s) - 1

	for tempLeft <= middle-1 && tempRight <= high {
		if helper[tempLeft].Weight <= helper[tempRight].Weight {
			s[current] = helper[tempLeft]
			tempLeft++
		} else {
			s[current] = helper[tempRight]
			tempRight++
		}
		current++
	}

	for tempLeft <= middle-1 {
		s[current] = helper[tempLeft]
		current++
		tempLeft++
	}
}