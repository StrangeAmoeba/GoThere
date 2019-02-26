package server

import (
  dt "concurrency-9/data_types"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

func Dist_matrix() {
  // for key, val := range Locations() {
  //   fmt.Printf("key[%s] value[%s]\n", k, v)
  // }
  // var url = constructURL(Locations()["sainagar"], Locations()["bhel"]) // external_api
  // fmt.Println(url) working
  // var content = getResponse(url)
  // fmt.Printf("%s\n", content) // for debugging
  getRespFile()
}

func getRespFile() {
  // Open our jsonFile
  // f, _ := os.OpenFile("server/notes.txt", os.O_RDWR|os.O_CREATE, 0755)
  // f.Close()
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

  fmt.Println(directions.Status)
}
