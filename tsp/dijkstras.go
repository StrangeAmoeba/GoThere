package tsp

import (
	"math"
	"sync"
)

// The maximum value of a float64 variable
const MaxFloat = math.MaxFloat64

// FindMinVertex is a helper function for the SingleSourceDijkstras function.
// It checks all the vertices in the graph and returns the index of the vertex
// with minimum weight.
//
//  Inputs: weights of individual vertices i.e. []float64,
//          tracker, that is used to track the visited vertices i.e. []bool
//  Output:  minIndex, the node of the graph with the least weight i.e. int
func FindMinVertex(weights []float64, tracker []bool) int {
	// TODO: Improve Concurrency of the function

	var min float64 = MaxFloat
	var minIndex int

	for i := range weights {
		if !tracker[i] && weights[i] <= min {
			min = weights[i]
			minIndex = i
		}
	}

	return minIndex
}

// SingleSourceDijkstras is a helper function for the dijkstras function
// The function returns the least weight for all paths from a single source vertex
// to all the other vertices in graph.
// Input: A [][]float64 slice, representing the Adjacency Matrix for the graph
//        A single int, representing the source vertex
// Output: A []float64 slice containg the smallest weights for all paths from the source
//         vertex to other vertices
func SingleSourceDijkstras(matrix [][]float64, src int) []float64 {
	var tracker []bool = make([]bool, len(matrix))
	var minWeights []float64 = make([]float64, len(matrix))

	for i := range tracker {
		tracker[i] = false
		minWeights[i] = MaxFloat
	}

	minWeights[src] = 0

	for i := 0; i < len(matrix)-1; i++ {
		var minVertex int = FindMinVertex(minWeights, tracker)

		tracker[minVertex] = true

		for j := range matrix {
			if !tracker[j] && matrix[minVertex][j] != 0 &&
				minWeights[minVertex] != MaxFloat &&
				minWeights[minVertex]+matrix[minVertex][j] < minWeights[j] {
				minWeights[j] = minWeights[minVertex] + matrix[minVertex][j]
			}
		}
	}

	return minWeights
}

// dijkstras parses a graph, and returns the least weight of all the paths between
// any two vertices in the graph
// Input: A [][]float64 slice, representing the Adjacency Matrix of the graph
// Output: A [][]float64 slice, which contains the minimum weight for the shortest path
//         between two vertices
func dijkstras(matrix [][]float64) [][]float64 {
	var minGraph [][]float64 = make([][]float64, len(matrix))
	var nodeWg sync.WaitGroup

	nodeWg.Add(len(matrix))

	for src := range matrix {
		go func(src int) {
			var minWeights []float64 = make([]float64, len(matrix))

			minWeights = SingleSourceDijkstras(matrix, src)

			minGraph[src] = minWeights

			nodeWg.Done()
		}(src)
	}

	nodeWg.Wait()

	return minGraph
}
