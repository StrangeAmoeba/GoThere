package server

import (
  dt "concurrency-9/data_types"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "math/rand"
  "os"
  "strings"
  "time"
)

// dist_traffic uses the google_directions_api and with the help of
// defined struct to parse json, we obtain distance and traffic weights
// calls assign_weight() to normalize the two parameters into one weight.
//
// Input: key1[ to assign the weight to the Dist_matrix ] i.e. int, loc1[ origin of the route ] i.e. string
// key1[ to assign the weight to the Dist_matrix ] i.e. int, loc2[ origin of the route ] i.e. string
// Output: weight[ weight of edge between the two locations ] i.e. float64
func dist_traffic(key1 int, loc1 string, key2 int, loc2 string) {
  var limit = false
  for {
    var url = constructURL(Locations()[loc1], Locations()[loc2]) // external_api

    var content = getResponse(url)

    var directions dt.Dir_info
    json.Unmarshal([]byte(content), &directions)

    if strings.Compare(directions.Status, "OVER_QUERY_LIMIT") == 0 {
      fmt.Println("OVER QUERY LIMIT")

      // now call with 1s timer, then we determine daily limit is exceeded
      if limit == false {
        time.Sleep(1100 * time.Millisecond)
        limit = true // now the error should not be met
        continue
      }

      // else if there is actually a limit, we fall back and load our backup data
      updated_matrix = true
      return
    } else if strings.Compare(directions.Status, "OK") != 0 {
      fmt.Println("ERROR - GOOGLE API REJECTED QUERY")
      updated_matrix = true
      return
    }

    var dist = directions.Routes[0].Legs[0].Distance.Val
    var traff = directions.Routes[0].Legs[0].Duration_traffic.Val
    // var dist = randFloats(1000.0, 10000.0, 1) // debugging
    // var traff = randFloats(1000.0, 6000.0, 1) // debugging

    Dist_matrix[key1][key2] = assign_weight(dist, traff) // distance_matrix
    return
    // getRespFile() - for debugging only
  }
}

// getRespFile is a helper function
// useful for debugging. Can access json file to obtain the parsed data.
// uses (server/example-route.json).
//
// Input: None
// Output: None
func getRespFile() {
  jsonFile, err := os.Open("server/example-route.json")

  // if we os.Open returns an error then handle it
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Successfully Opened json file")

  // defer the closing of our jsonFile so that we can parse it later on
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  var directions dt.Dir_info
  json.Unmarshal([]byte(byteValue), &directions)
}

// randFloats is a helper function
// useful for debugging. Generate an array of random float64
//
// Input: min[ minimum bound ] i.e. float64, max[ maximum bound ] i.e. float64
// n[ size of random numbers array ] i.e. int
// Output: res[ an array of random numbers ] i.e. []float64
func randFloats(min, max float64, n int) []float64 {
  res := make([]float64, n)

  for i := range res {
    res[i] = min + rand.Float64()*(max-min)
  }

  return res
}
