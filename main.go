package main

import (
  "fmt"
  "concurrency-9/tsp"
)

func main() {
  matrix := tsp.Get_mat()
  a := tsp.Get_best_path(matrix, []int{3,4,7,22,25,48})
  fmt.Printf("%v", a)
}