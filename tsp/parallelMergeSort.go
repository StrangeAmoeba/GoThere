package tsp

import (
	"concurrency-9/dataTypes"
	"math"
	"sync"
)

// MergeSort, a parallel version of mergesort which brings down running time to O(n).
//
//  Input: s i.e. dataTypes.GraphEdge
//  Output: s, sorted i.e. dataTypes.GraphEdge
func MergeSort(s []dataTypes.GraphEdge) []dataTypes.GraphEdge {

	// nearest power of 2 to len(s)
	var k float64 = math.Ceil(math.Log2(float64(len(s))))

	// number of numbers to be added to the slice
	var diff int = int(math.Exp2(k) - float64(len(s)))

	// looping to add the deficit
	for i := 0; i < diff; i++ {
		s = append(s, dataTypes.GraphEdge{-1, -1, -1})
	}

	// call normal mergesort for smaller arrays
	if len(s) < 1024 {
		normalMergesort(s)
	} else {
		parallelMergesort(s)
	}

	return s[diff:]
	// fmt.Printf("%v", s)
}

func normalMergesort(s []dataTypes.GraphEdge) {

	// straight-forward implementation of mergesort

	if len(s) > 1 {
		middle := len(s) / 2
		// Split in Middle and continue
		normalMergesort(s[:middle])
		normalMergesort(s[middle:])
		merge(s, middle)
	}
}

func parallelMergesort(s []dataTypes.GraphEdge) {
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
			parallelMergesort(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergesort(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

func merge(s []dataTypes.GraphEdge, middle int) {
	helper := make([]dataTypes.GraphEdge, len(s))
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
