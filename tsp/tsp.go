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

	// var a []int  = []int{4,5,6,7,8}
	// a = append(a[:2], a[3:]...)
	// fmt.Printf("%v",a)
	// fmt.Printf("%v", create_destination_matrix(matrix, []int{3,4,7,25,48}))
	return kruskals(create_destination_matrix(matrix, []int{3,4,7,25,48}))
	// fmt.Printf("%v", Merge_Sort(s))
}

func create_destination_matrix(matrix[][] float64, destinations []int) [][]float64{
	var temp, temp1 [][]float64
	// fmt.Printf("%d", len(destinations))
	for _, i := range destinations {
		temp = append(temp, matrix[i])
	}
	for i:=0;i<len(destinations);i++ {
		// fmt.Printf("%d", i)
		var row []float64
		// temp = append(temp, row)
		for _,j := range destinations {
			// fmt.Printf("%d", j)
			row = append(row, temp[i][j])
		}
		temp1 = append(temp1, row)
	}
	// fmt.Printf("%v", temp)
	return temp1
}

func Get_best_path(matrix [][]float64, destinations []int) {

	destinations_matrix := create_destination_matrix(matrix, destinations)
	destinations_matrix = destinations_matrix
	// mst := kruskals(destinations_matrix)
	// best_path := get_best_path_from_mst
	//return path
}