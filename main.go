package main

import (
  "concurrency-9/server"
  "concurrency-9/tsp"
  "fmt"
  "log"
  "net/http"
  "os"
  "sort"
  "strings"
)

// getIndices is responsible to parse through the form response given by form.html to
// find the user queried locations. The parsed data will consist of locations which
// in turn will be converted to indices, each representing their index in the DistMatrix.
//
// Input: loc [ user queried locations from the form ] i.e. map[string][]string
// Output: indices [ array of user queries locations in indices ] i.e. []int
func getIndices(loc map[string][]string) []int {
  var count = len(loc)
  count = count / 2 // we dont need the key value of the field. only its value suffices

  var indices = make([]int, 1, 1)
  var locKeyRaw = loc["form_data[0][value]"][0]
  locKeyRaw = strings.ToLower(locKeyRaw)
  result := strings.Split(locKeyRaw, " ")
  var length = len(result)

  var locKey strings.Builder
  for i := 0; i < length; i++ {
    fmt.Fprintf(&locKey, result[i])
  }

  indices[0] = server.Locations()[locKey.String()].Index

  for i := 2; i <= count; i++ {
    var key strings.Builder
    fmt.Fprintf(&key, "form_data[%d][value]", i-1)
    locKeyRaw = loc[key.String()][0]
    locKeyRaw = strings.ToLower(locKeyRaw)

    result = strings.Split(locKeyRaw, " ")
    length = len(result)
    locKey.Reset()

    for i := 0; i < length; i++ {
      fmt.Fprintf(&locKey, result[i])
    }

    var locInd = server.Locations()[locKey.String()].Index
    indices = append(indices, locInd)
  }

  return indices
}

// determineListenAddress figures out what address to listen on for traffic.
// It uses the $PORT environment variable only to determine this.
// If $PORT isn’t set an error is returned instead.
//
// Input: none
// Output: port[ $PORT env variable ] i.e. string, err[ $PORT not set ] i.e. error
func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return ":9000", fmt.Errorf("$PORT not set, using :9000")
  }
  return ":" + port, nil
}

// serveForm is a handler which responds to an HTTP request.
// Currently supports GET and POST requests.
// Serves form.html in public
//
// Input: w [ used to construct an HTTP response. ] i.e. http.ResponseWriter,
// r [ pointer to http Request ] i.e. *http.Request
// Output: None
func serveForm(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.Error(w, "404 not found.", http.StatusNotFound)
    return
  }

  switch r.Method {
  case "GET":
    http.ServeFile(w, r, "public/form.html")

  case "POST":
    if err := r.ParseForm(); err != nil {
      fmt.Fprintf(w, "ParseForm() err: %v", err)
      return
    }

    var indices = getIndices(r.Form)

    sort.Ints(indices) // sort the locations indices in increasing order
    var bestPath, route_helper = tsp.GetBestPath(server.DistSliceMatrix, indices)

    // keys to get the locations - Lat and Long, from given index
    keys := make([]string, 0)
    for k := range server.Locations() {
      keys = append(keys, k)
    }
    // give ordering for optimizing search
    sort.Strings(keys)

    // best_path markers
    var length = len(bestPath)

    // write with JSON parsable string syntax
    var json strings.Builder
    // start with json stringified array and enter first location
    fmt.Fprintf(&json, "{\"path\":[[\"%v\", \"%v\", \"%v\"]",
      server.LocKeys()[bestPath[0]],
      server.Locations()[keys[bestPath[0]]].Lat,
      server.Locations()[keys[bestPath[0]]].Long)

    // append the locations to the json stringified array
    for i := 1; i < length; i++ {
      fmt.Fprintf(&json, ", [\"%v\", \"%v\", \"%v\"]",
        server.LocKeys()[bestPath[i]],
        server.Locations()[keys[bestPath[i]]].Lat,
        server.Locations()[keys[bestPath[i]]].Long)
    }
    // close the json stringified array for the best path
    fmt.Fprintf(&json, "], ")

    // route_markers
    length = len(route_helper)

    // continue with JSON parsable string syntax
    // start route with json stringified array and the first route marker if it exists
    fmt.Fprintf(&json, "\"route\":[")
    if length != 0 {
      fmt.Fprintf(&json, "[\"%v\", \"%v\", \"%v\"]",
        server.LocKeys()[route_helper[0]],
        server.Locations()[keys[route_helper[0]]].Lat,
        server.Locations()[keys[route_helper[0]]].Long)
    }

    // append the locations to the json stringified array
    for i := 1; i < length; i++ {
      fmt.Fprintf(&json, ", [\"%v\", \"%v\", \"%v\"]",
        server.LocKeys()[route_helper[i]],
        server.Locations()[keys[route_helper[i]]].Lat,
        server.Locations()[keys[route_helper[i]]].Long)
    }
    // close the json stringified array
    fmt.Fprintf(&json, "]}")

    // We write an object in strigified form where -
    // each array in the array consists of info, lat, long of a place respectively.
    fmt.Fprintf(w, "%v", json.String())
    json.Reset()

  default:
    fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
  }
}

func main() {
  // testing - harsha
  server.CreateDistMatrix()

  // web app
  addr, err := determineListenAddress()
  if err != nil {
    fmt.Println(err)
  }

  http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
  http.HandleFunc("/", serveForm)

  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
