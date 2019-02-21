package tsp

import (
	dt "concurrency-9/data_types"
	// "fmt"
	"math"
)
const number_of_nodes = 50
var float_max = math.MaxFloat64

func Getmat() [][] float64{
	// currently just dummy matrix
	var matrix [][] float64
	for i:=0;i<number_of_nodes;i++ {
		temp := make([]float64, number_of_nodes)
		for j:=0;j<number_of_nodes;j++ {
			temp[j] = float64((j-i)*(j-i))
		}
		matrix = append(matrix,temp)
	}

	return matrix
}

func Get_MST(matrix [][]float64) []dt.Graph_edge {

	return kruskals(matrix)
	// fmt.Printf("%v", Merge_Sort(s))
}