package server

import (
  "fmt"
  "sort"
  "sync"
)

var Dist_matrix = [50][50]float64{}
var avg_speed = 15.5 // mps

// Create_dist_matrix is responsible for assigning weights between each location
// which is important for the application of dijkstas and kruskals algorithm to find the shortest
// route. The matrix created can later be accessed using Dist_matrix which is of the type [][]float64
// We use go routines to create the matrix as quickly as possible.
// WaitGroup is used to find if all go routines have completed their execution.
// Input: None
// Output: None
func Create_dist_matrix() {
  var wg sync.WaitGroup
  keys := make([]string, 0)
  for k_i := range Locations() {
    keys = append(keys, k_i)
  }
  sort.Strings(keys)
  for k_i, v_i := range keys {
    wg.Add(1)
    go func(k_i int, v_i string, keys []string) {
      defer wg.Done() // Decrement the counter when the goroutine completes.
      for k_j, v_j := range keys {
        if k_i != k_j {
          fmt.Println(k_i, v_i, k_j, v_j)
          dist_traffic(k_i, v_i, k_j, v_j) // json_parse_dist_traff - assign values to the matrix
        }
      }
    }(k_i, v_i, keys)
  }
  // Wait for all go routines to complete
  wg.Wait()
}

// assign_weight is responsible to normalize the two weights - distance and traffic
// into one weight.
// Input: dist[ weight representing distance of route ] i.e. float64,
// traff[ weight representing traffic in route ] i.e. float64
// Output: weight[ weight of edge between the two locations ] i.e. float64
func assign_weight(dist, traff float64) float64 {
  var weight = dist
  weight += (dist * avg_speed)
  return weight
}
