package main

import (
  "fmt"
  "concurrency-9/tsp"
)

func main() {
  matrix := tsp.Get_mat()
  a := tsp.Get_best_path(matrix, []int{3,4,7,25,48})
  // a = a
  fmt.Printf("%v", a)
  // tsp.Get_adjacency_list()
}