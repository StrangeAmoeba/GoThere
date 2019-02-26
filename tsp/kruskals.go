package tsp

import (
	// "fmt"
	dt "concurrency-9/data_types"
	"sync"
)

// this will take input the matrix coming in from the API
// and convert it into an adjacency list (input type matrix [][]float64)

var number_of_vertices int

// Get_adjacency_list converts a graph in the form of an adjacency matrix to
// an adjacency list.
//
//   Input: matrix i.e. [][]float64
//   Output: al i.e. [][]data_types.Graph_edge
func Get_adjacency_list(matrix [][]float64) [][]dt.Graph_edge {
	var al [][]dt.Graph_edge
	for i := 0; i < len(matrix); i++ {
		temp := make([]dt.Graph_edge, number_of_vertices)
		for j := 0; j < len(matrix[i]); j++ {
			temp[j] = dt.Graph_edge{i, j, matrix[i][j]}
		}
		al = append(al, temp)
	}
	// fmt.Printf("%v", al[0])
	return al
}

type subset struct {
	parent, rank int
}

func find(subsets []subset, i int) int {
	if subsets[i].parent != i {
		subsets[i].parent = find(subsets, subsets[i].parent)
	}

	return subsets[i].parent
}

// Union computes the union of two disjoint sets, this particular implementation
// uses parallelisation of any adjacent calls to improve speed.
//
//   Input: subsets, the list of all subsets, i.e. []subset
//   		 x, y,representative elements of subsets, i.e. int.
//   Output: No Output, union stored in subsets.
func Union(subsets []subset, x int, y int) {

	// concurrently find the roots of both the
	var wg sync.WaitGroup
	wg.Add(2)
	var xroot, yroot int
	go func() {
		defer wg.Done()
		xroot = find(subsets, x)
	}()
	go func() {
		defer wg.Done()
		yroot = find(subsets, y)
	}()
	wg.Wait()
	// Attach smaller rank tree under root of high
	// rank tree (Union by Rank)
	if subsets[xroot].rank < subsets[yroot].rank {
		subsets[xroot].parent = yroot
	} else if subsets[xroot].rank > subsets[yroot].rank {
		subsets[yroot].parent = xroot
	} else {
		subsets[yroot].parent = xroot
		subsets[xroot].rank++
	}
}

// Kruskals computes the mst of a given graph represented in the form of a an adjacency matrix
// , internally it is converted to an adjacency list. This is a parallelised form of Kruskal's algorithm,
// where any adjacent calls are made parallely.
//
//  Input: matrix i.e. [][]float64.
//  Output: results i.e. data_types.Graph_edge.
func Kruskals(matrix [][]float64) []dt.Graph_edge {
	number_of_vertices = len(matrix)
	V := number_of_vertices
	var results []dt.Graph_edge
	s := Get_adjacency_list(matrix)
	e := 0 // number of edges in results
	var edges []dt.Graph_edge
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s[i]); j++ {
			edges = append(edges, dt.Graph_edge{i, j, (s[i][j].Weight + s[j][i].Weight) / 2.0})
		}
	}
	edges = Merge_Sort(edges)
	var subsets []subset
	for i := 0; i < V; i++ {
		subsets = append(subsets, subset{i, 0})
	}
	i := 0
	var wg sync.WaitGroup
	for e < V-1 {
		next_edge := edges[i]
		i++
		var x, y int

		// parallely find the sets of x and y

		wg.Add(2)
		go func() {
			defer wg.Done()
			x = find(subsets, next_edge.Src)
		}()
		go func() {
			defer wg.Done()
			y = find(subsets, next_edge.Dst)
		}()
		wg.Wait()

		if x != y {
			results = append(results, next_edge)
			e++
			Union(subsets, x, y)
		}
	}
	return results
}
