package main

import (
  "concurrency-9/server"
  // "concurrency-9/tsp"
  "fmt"
  "log"
  "net/http"
  "os"
  // "sort"
  "strings"
)

// get_indices is responsible to parse through the form response given by form.html to
// find the user queried locations. The parsed data will consist of locations which
// in turn will be converted to indices, each representing their index in the Dist_matrix.
// Input: loc [ user queried locations from the form ] i.e. map[string][]string
// Output: indices [ array of user queries locations in indices ] i.e. []int
func get_indices(loc map[string][]string) []int {
  var count = len(loc)
  var indices = make([]int, 1, 1)
  var loc_key_raw = loc["loc1"][0]
  loc_key_raw = strings.ToLower(loc_key_raw)
  result := strings.Split(loc_key_raw, " ")
  var length = len(result)
  var loc_key strings.Builder
  for i := 0; i < length; i++ {
    fmt.Fprintf(&loc_key, result[i])
  }
  indices[0] = server.Locations()[loc_key.String()].Index
  for i := 2; i <= count; i++ {
    var key strings.Builder
    fmt.Fprintf(&key, "loc%d", i)
    loc_key_raw = loc[key.String()][0]
    loc_key_raw = strings.ToLower(loc_key_raw)
    result = strings.Split(loc_key_raw, " ")
    length = len(result)
    loc_key.Reset()
    for i := 0; i < length; i++ {
      fmt.Fprintf(&loc_key, result[i])
    }
    var loc_ind = server.Locations()[loc_key.String()].Index
    indices = append(indices, loc_ind)
  }
  return indices
}

// determineListenAddress figures out what address to listen on for traffic.
// It uses the $PORT environment variable only to determine this.
// If $PORT isn’t set an error is returned instead.
// Input: none
// Output: port[ $PORT env variable ] i.e. string, err[ $PORT not set ] i.e. error
func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

// serveForm is a handler which responds to an HTTP request.
// Currently supports GET and POST requests.
// Serves form.html in public
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
    fmt.Println(r.Form)
    // var indices = get_indices(r.Form)
    // sort.Ints(indices) // sort the locations indices in increasing order
    // var dist_slice_matrix = server.MatToDynMat()
    // best_path := tsp.Get_best_path(dist_slice_matrix, indices)

    // for i := 0; i < len(best_path); i++ {
    //   for j := range server.Locations() {
    //     if best_path[i] == server.Locations()[j].Index {
    //       fmt.Println(best_path[i], j)
    //       break
    //     }
    //   }
    // }
  default:
    fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
  }
}

func main() {
  // testing - harsha
  server.Create_dist_matrix()

  // web app
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }

  http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
  http.HandleFunc("/", serveForm)

  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
