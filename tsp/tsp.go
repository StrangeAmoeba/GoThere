package tsp

import (
	dt "concurrency-9/dataTypes"
	"fmt"
	"math"
)

const numberOfNodes = 35

var floatMax = math.MaxFloat64

// Getmat returns a dummy matrix of dimension mentioned in the file tsp.go
//
// Output: matrix, dummy matrix, i.e. [][]float64
func GetMat() [][]float64 {
	// currently just dummy matrix
	var matrix [][]float64
	for i := 0; i < numberOfNodes; i++ {
		temp := make([]float64, numberOfNodes)
		for j := 0; j < numberOfNodes; j++ {
			temp[j] = float64(1.0 / (math.Abs(float64(j - i))))
		}
		matrix = append(matrix, temp)
	}

	return matrix
}

// GetMST is a wrapper around destinationsMatrix() and kruskals()
// it calls these two and returns the required matrix for tsp computation
//
// Input: matrix i.e. [][]float64, destinations i.e. int
// Output: mst i.e. minimum spanning tree of the subgraph, i.e. []dataTypes.GraphEdge
func GetMST(matrix [][]float64, destinations []int) []dt.GraphEdge {

	destinationsMatrix := CreateDestinationMatrix(matrix, destinations)
	mst := Kruskals(destinationsMatrix)
	return mst
}

// CreateDestinationMatrix takes in the list of destinations and creates a graph
// in the form of an adjacency matrix with just these nodes and and edges containing them
// from the parent graph, there is no loss of data in this as whenever it is called,it is after
// an algorithm (say, dijkstra's) which incapsulates any important data from other nodes
//
// Input: matrix, the parent graph i.e. [][]float64, destinations i.e. []int
// Output: destinations i.e. [][]float64
func CreateDestinationMatrix(matrix [][]float64, destinations []int) [][]float64 {
	var temp, temp1 [][]float64
	for _, i := range destinations {
		temp = append(temp, matrix[i])
	}
	for i := 0; i < len(destinations); i++ {
		var row []float64
		for _, j := range destinations {
			row = append(row, temp[i][j])
		}
		temp1 = append(temp1, row)
	}
	return temp1
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetBestPath, takes in a graph, a subset of nodes of the same graph and gives the ideal
// path to visit all of them in the least expensive way.
//
// Input: matrix, the full graph i.e. [][]float64, destinations, subset of nodes i.e. []int
// Output: bestPath i.e. []int
func GetBestPath(matrix [][]float64, destinations []int) []int {
	// fmt.Printf("%v", destinations)
	matrix, internalNodes := Dijkstras(matrix)
	// fmt.Println(internalNodes)
	mst := GetMST(matrix, destinations)
	// makeTreeFromEdges(mst, len(destinations))
	pw := preOrderWalk(mst)
	var bestPath []int
	for i := 0; i < len(pw); i++ {
		bestPath = append(bestPath, destinations[pw[i]])
	}
	var bestPathWithInternalNodes []int
	for i := 0; i < len(bestPath)-1; i++ {
		bestPathWithInternalNodes = append(bestPathWithInternalNodes, internalNodes[bestPath[i]][bestPath[i+1]]...)
		if i != len(bestPath)-2 {
			bestPathWithInternalNodes = bestPathWithInternalNodes[0 : len(bestPathWithInternalNodes)-1]
		}
	}
	fmt.Println(bestPath, unique(bestPathWithInternalNodes))
	return bestPathWithInternalNodes
}
