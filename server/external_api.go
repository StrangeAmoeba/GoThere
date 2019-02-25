package server

import (
  dt "concurrency-9/data_types"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

// constructURL takes the origin and destination Vertices which consists of
// its latitude and longitude. We contruct a url with accordance to google api requirement
// for parsing the json file which we will receive later on.
// Input: origin [ origin vertex ] i.e. dt.Vertex, destination [ destination vertex ] i.e. dt.Vertex
// Output: url [ url link address ] i.e. string
func constructURL(origin, dest dt.Vertex) string {
  var url strings.Builder
  fmt.Fprintf(&url, "https://maps.googleapis.com/maps/api/directions/json?origin=")
  // server key
  var key = "AIzaSyDUJjTw3LKudLIxSj_saJQLHsmH_RTfa9w"
  // attach origin co-ordinates to the url
  fmt.Fprintf(&url, "%s%s%s%s", fmt.Sprint(origin.Lat), ",", fmt.Sprint(origin.Long), "&destination=")
  // attach origin co-ordinates to the url
  fmt.Fprintf(&url, "%s%s%s", fmt.Sprint(dest.Lat), ",", fmt.Sprint(dest.Long))
  // attach the key to the url
  fmt.Fprintf(&url, "%s%s%s", "&key=", key, "&departure_time=now&traffic_model=pessimistic")
  return url.String()
}

// getResponse is responsible for fetching the content of website which is specified
// via the url.
// Input: url [ url link address ] i.e. string
// Output: html [ html content in bytes ] i.e. []byte
func getResponse(url string) []byte {
  resp, err := http.Get(url)
  // handle the error if there is one
  if err != nil {
    panic(err)
  }
  // do this now so it won't be forgotten
  defer resp.Body.Close()
  // reads html as a slice of bytes
  html, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  return html
}
