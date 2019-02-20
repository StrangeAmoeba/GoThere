package tsp

import (
	"fmt"
	"concurrency-9/data_types"
)

// this will tak input the matrix coming in from the API
// and convert it into an adjacency list (input type matrix [][]float64)

func Get_adjacency_list () [][]data_types.Graph_edge {
	var s [][] data_types.Graph_edge
	matrix := getmat()
	for i:=0;i<len(matrix);i++ {
		temp := make([]data_types.Graph_edge, number_of_nodes)
		for j:=0;j<len(matrix[i]);j++ {
			temp[j] = data_types.Graph_edge{i,j,matrix[i][j]}
		}
		s = append(s, temp)
	}
	fmt.Printf("%v", s[0])
	return s
}

func kruskals ( adjacency_list [][]data_types.Graph_edge) {
	var set []int = []int {0}
	set = set
}