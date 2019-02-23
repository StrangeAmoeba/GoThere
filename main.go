package main

import (
  "concurrency-9/server"
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
  fmt.Fprintln(w, "Hello World!!")
}

func main() {
  // testing - harsha
  server.Create_dist_matrix()

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
