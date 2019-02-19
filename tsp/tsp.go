package tsp

import (
	"concurrency-9/data_types"
	// "fmt"
	"math"
)
const number_of_nodes = 50
var float_max = math.MaxFloat64

func getmat() [][] data_types.Weight_tuple {
	// currently just dummy matrix
	var matrix [][] data_types.Weight_tuple
	for i:=0;i<number_of_nodes;i++ {
		temp := make([]data_types.Weight_tuple, number_of_nodes)
		for j:=0;j<number_of_nodes;j++ {
			temp[j] = data_types.Weight_tuple{float64(i),float64(j)}
		}
		matrix = append(matrix,temp)
	}

	return matrix
}

func Get_MST() []float64 {

	var s = []float64{4,5,8,15,6}
	s = Merge_Sort(s)
	return s
	// fmr.Printf("%v", Merge_Sort(s))
}