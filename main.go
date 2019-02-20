package main

import (
  "fmt"
  "concurrency-9/tsp"
)

func main() {
  a := tsp.Get_MST()
  fmt.Printf("%v", a)
  // tsp.Get_adjacency_list()
}