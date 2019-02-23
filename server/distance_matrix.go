package server

import (
  dt "concurrency-9/data_types"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "sort"
  "strings"
)

var Dist_matrix = [50][50]float64{}

// Create_dist_matrix is responsible for assigning weights between each location
// which is important for the application of dijkstas and kruskals algorithm to find the shortest
// route. The matrix created can later be accessed using Dist_matrix which is of the type [][]float64
// Input: None
// Output: None
func Create_dist_matrix() {
  keys := make([]string, 0)
  for k_i := range Locations() {
    keys = append(keys, k_i)
  }
  sort.Strings(keys)
  for _, v_i := range keys {
    for _, v_j := range keys {
      if strings.Compare(v_i, v_j) != 1 {
        go dist_traffic(v_i, v_j) // json_parse_dist_traff - assign values to the matrix
      }
    }
  }
}

// assign_weight is responsible to normalize the two weights - distance and traffic
// into one weight.
// Input: dist[ weight representing distance of route ] i.e. int,
// traff[ weight representing traffic in route ] i.e. int
// Output: weight[ weight of edge between the two locations ] i.e. float64
func assign_weight(dist, traff int) float64 {

}
