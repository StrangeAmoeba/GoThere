package server

import (
  "fmt"
  "net/http"
  "io/ioutil"
  dt "concurrency-9/data_types"
  )

func constructURL(origin, dest dt.Vertex) string {
  var url = "https://maps.googleapis.com/maps/api/directions/json?origin="
  // server key
  var key = "AIzaSyDUJjTw3LKudLIxSj_saJQLHsmH_RTfa9w"
  // attach origin co-ordinates to the url
  url += fmt.Sprint(origin.Lat) + "," + fmt.Sprint(origin.Long) + "&destination="
  // attach origin co-ordinates to the url
  url += fmt.Sprint(dest.Lat) + "," + fmt.Sprint(dest.Long)
  // attach the key to the url
  url += "&key=" + key
  return url
}

func getResponse(url string) []byte{
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
