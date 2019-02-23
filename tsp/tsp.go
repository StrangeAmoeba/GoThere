package tsp

import (
	dt "concurrency-9/data_types"
	// "fmt"
	"math"
)
const number_of_nodes = 50
var float_max = math.MaxFloat64

//Getmat returns a dummy matrix of dimension mentioned in the file tsp.go
//
//  Output: matrix, dummy matrix, i.e. [][]float64
func Get_mat() [][] float64{
	// currently just dummy matrix
	var matrix [][] float64
	for i:=0;i<number_of_nodes;i++ {
		temp := make([]float64, number_of_nodes)
		for j:=0;j<number_of_nodes;j++ {
			temp[j] = float64(1.0/(math.Abs(float64(j-i))))
		}
		matrix = append(matrix,temp)
	}

	return matrix
}

// Get_MST is a wrapper around destinations_matrix() and kruskals()
// it calls these two and returns the required matrix for tsp computation
// 
//  Input: matrix i.e. [][]float64, destinations i.e. int
//  Output: mst i.e. minimum spanning tree of the subgraph, i.e. []data_types.Graph_edge
func Get_MST(matrix [][]float64, destinations []int) []dt.Graph_edge {

	destinations_matrix := Create_destination_matrix(matrix, destinations)
	mst := Kruskals(destinations_matrix)
	return mst
}


// Create_destination_matrix takes in the list of destinations and creates a graph
// in the form of an adjacency matrix with just these nodes and and edges containing them
// from the parent graph, there is no loss of data in this as whenever it is called,it is after
// an algorithm (say, dijkstra's) which incapsulates any important data from other nodes
//
//  Input: matrix, the parent graph i.e. [][]float64, destinations i.e. []int
//
//  Output: destinations i.e. [][]float64
func Create_destination_matrix(matrix[][] float64, destinations []int) [][]float64{
	var temp, temp1 [][]float64
	for _, i := range destinations {
		temp = append(temp, matrix[i])
	}
	for i:=0;i<len(destinations);i++ {
		var row []float64
		for _,j := range destinations {
			row = append(row, temp[i][j])
		}
		temp1 = append(temp1, row)
	}
	return temp1
}


// Get_best_path, takes in a graph, a subset of nodes of the same graph and gives the ideal
// path to visit all of them in the least expensive way.
//
//  Input: matrix, the full graph i.e. [][]float64, destinations, subset of nodes i.e. []int
//  Output: best_path i.e. []int
func Get_best_path(matrix [][]float64, destinations []int) []int {
	// fmt.Printf("%v", destinations)
	matrix = dijkstras(matrix)
	mst := Get_MST(matrix, destinations)
	// make_tree_from_edges(mst, len(destinations))
	pw := pre_order_walk(mst)
	var best_path []int
	for i:=0;i<len(pw);i++ {
		best_path = append(best_path, destinations[pw[i]])
	}
	return best_path
}