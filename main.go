package main

import (
  // "concurrency-9/server"
  "fmt"
  "log"
  "net/http"
  "os"
)

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

func hello(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.Error(w, "404 not found.", http.StatusNotFound)
    return
  }

  switch r.Method {
  case "GET":
    http.ServeFile(w, r, "form.html")
  case "POST":
    if err := r.ParseForm(); err != nil {
      fmt.Fprintf(w, "ParseForm() err: %v", err)
      return
    }
    fmt.Fprintf(w, "%v\n", r.PostForm["locations"])
    name := r.FormValue("name")
    address := r.FormValue("noou")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "noou = %s\n", address)
  default:
    fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
  }
}

func main() {
  // testing - harsha
  // server.Create_dist_matrix()

  // web app
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/", hello)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
