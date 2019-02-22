package server

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
  dt "concurrency-9/data_types"
  )

func Dist_matrix() {
  var url = constructURL(Locations()["sainagar"], Locations()["bhel"]) // external_api
  fmt.Println(url)
  var content = getResponse(url)
  var directions dt.Dir_info
  json.Unmarshal([]byte(content), &directions)
  fmt.Println("debug3", directions.Routes[0].Legs[0].Distance.Val)
  // getRespFile() - for debugging only
}

// helper function - useful for debugging. Can access json files.
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
