package server

// import (
//   "fmt"
// )

func Dist_matrix() {
  // for key, val := range Locations() {
  //   fmt.Printf("key[%s] value[%s]\n", k, v)
  // }
  var url = constructURL(Locations()["sainagar"], Locations()["bhel"]) // external_api
  // fmt.Println(url) working
  getResponse(url)
  // fmt.Printf(content)
}
