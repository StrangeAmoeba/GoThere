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
// Inputs: Weights of individual vertices, as a []float64 slice
//         A []bool slice, that is used to track the visited vertices
// Output: A single integer representing the node of the graph with the least weight
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
//
// Input:   A [][]float64 slice, representing the Adjacency Matrix for the graph
//          A single int, representing the source vertex
// Outputs: A []float64 slice containg the smallest weights for all paths from the source
//          vertex to other vertices
//          A 2D slice containing the minimum paths from vertex i to j i.e. [][]int
func SingleSourceDijkstras(matrix [][]float64, src int) ([]float64, [][]int) {
	var tracker []bool = make([]bool, len(matrix))
	var minWeights []float64 = make([]float64, len(matrix))
	var minPath [][]int = make([][]int, len(matrix))

	for i := range tracker {
		tracker[i] = false
		minWeights[i] = MaxFloat
	}

	minWeights[src] = 0
	minPath[src] = append(minPath[src], src)

	for i := 0; i < len(matrix)-1; i++ {
		var minVertex int = FindMinVertex(minWeights, tracker)

		tracker[minVertex] = true

		for j := range matrix {
			if !tracker[j] && matrix[minVertex][j] != 0 &&
				minWeights[minVertex] != MaxFloat &&
				minWeights[minVertex]+matrix[minVertex][j] < minWeights[j] {
				minWeights[j] = minWeights[minVertex] + matrix[minVertex][j]
				minPath[j] = append(minPath[minVertex], j)
			}
		}
	}

	return minWeights, minPath
}

// Dijkstras parses a graph, and returns the least weight of all the paths between
// any two vertices in the graph
//
// Input:   A [][]float64 slice, representing the Adjacency Matrix of the graph
// Outputs: A [][]float64 slice, which contains the minimum weight for the shortest path
//          between two vertices
//          A 3D int slice, representing the mininimum paths b/w any two vertices of the
//          graph i.e. [][][]int
func Dijkstras(matrix [][]float64) ([][]float64, [][][]int) {
	var minGraph [][]float64 = make([][]float64, len(matrix))
	var minPaths [][][]int = make([][][]int, len(matrix))
	var nodeWg sync.WaitGroup

	nodeWg.Add(len(matrix))

	for src := range matrix {
		go func(src int) {
			var minWeights []float64 = make([]float64, len(matrix))
			var minPath [][]int = make([][]int, len(matrix))

			minWeights, minPath = SingleSourceDijkstras(matrix, src)

			minGraph[src] = minWeights
			minPaths[src] = minPath

			nodeWg.Done()
		}(src)
	}

	nodeWg.Wait()

	return minGraph, minPaths
}
