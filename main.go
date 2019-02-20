package main

import (
  "fmt"
  "concurrency-9/tsp"
)

func main() {
  matrix := tsp.Getmat()
  a := tsp.Get_MST(matrix)
  fmt.Printf("%v", a)
  // tsp.Get_adjacency_list()
}