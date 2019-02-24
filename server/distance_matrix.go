package server

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "sort"
  "strconv"
  "strings"
  "sync"
  "time"
)

var Dist_matrix = [35][35]float64{}
var avg_speed = 15.5 // mps
var updated_matrix = false

// Create_dist_matrix is responsible for assigning weights between each location
// which is important for the application of dijkstas and kruskals algorithm to find the shortest
// route. The matrix created can later be accessed using Dist_matrix which is of the type [][]float64
// We use go routines to create the matrix as quickly as possible.
// WaitGroup is used to find if all go routines have completed their execution.
// Input: None
// Output: None
func Create_dist_matrix() {
  check_matrix_file()
  if updated_matrix == true {
    return
  }
  var wg sync.WaitGroup
  keys := make([]string, 0)
  for k_i := range Locations() {
    keys = append(keys, k_i)
  }
  sort.Strings(keys)
  for k_i, v_i := range keys {
    wg.Add(1)
    go func(k_i int, v_i string, keys []string) {
      defer wg.Done() // Decrement the counter when the goroutine completes.
      for k_j, v_j := range keys {
        if k_i != k_j {
          dist_traffic(k_i, v_i, k_j, v_j) // json_parse_dist_traff - assign values to the matrix
        }
        if k_j == 15 {
          time.Sleep(1000 * time.Millisecond)
        }
        if k_j == 25 {
          time.Sleep(1000 * time.Millisecond)
        }
      }
    }(k_i, v_i, keys)
  }
  // Wait for all go routines to complete
  wg.Wait()
  write_matrix_file()
}

// assign_weight is responsible to normalize the two weights - distance and traffic
// into one weight.
// Input: dist[ weight representing distance of route ] i.e. float64,
// traff[ weight representing traffic in route ] i.e. float64
// Output: weight[ weight of edge between the two locations ] i.e. float64
func assign_weight(dist, traff float64) float64 {
  var weight = dist
  weight += (dist * avg_speed)
  return weight / 1000
}

// check_matrix_file is responsible to check if the current date stamp on file
// is different from current date.(Pacific Time Zone)
// If the current date stamp matches with current date, then update Dist_matrix
// else update updated_matrix accordingly
// Input: None
// Output: None
func check_matrix_file() {
  file, err := os.Open("dist_matrix.log")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  // PST date to be verified
  IST, err := time.LoadLocation("Asia/Kolkata")
  if err != nil {
    fmt.Println(err)
    return
  }
  PST, err := time.LoadLocation("America/Los_Angeles")
  if err != nil {
    fmt.Println(err)
    return
  }

  const longForm = "2006-01-02 15:04:05"
  currentTime := time.Now()
  t, err := time.ParseInLocation(longForm, currentTime.Format("2006-01-02 15:04:05"), IST)
  result := strings.Split(t.In(PST).String(), " ")
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  if strings.Compare(scanner.Text(), result[0]) != 0 { // not-updated
    return
  }
  // updated, so read matrix from file
  updated_matrix = true
  for i := 0; i < 35; i++ {
    for j := 0; j < 35; j++ {
      var data = scanner.Text()
      if val, err := strconv.ParseFloat(data, 64); err == nil {
        Dist_matrix[i][j] = val
      }
    }
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

// write_matrix_file is responsible to store the generated weight matrix into a file for accessing
// the matrix on the same day later again. Saves api requests made to google api.
// Only writes if the current date stamp on file is different from current date.(Pacific Time Zone)
// Input: None
// Output: None
func write_matrix_file() {
  file, err := os.Create("dist_matrix.log")
  if err != nil {
    fmt.Println(err)
    file.Close()
    os.Exit(3)
  }
  // PST date to be entered
  IST, err := time.LoadLocation("Asia/Kolkata")
  if err != nil {
    fmt.Println(err)
    return
  }
  PST, err := time.LoadLocation("America/Los_Angeles")
  if err != nil {
    fmt.Println(err)
    return
  }

  const longForm = "2006-01-02 15:04:05"
  currentTime := time.Now()
  t, err := time.ParseInLocation(longForm, currentTime.Format("2006-01-02 15:04:05"), IST)
  result := strings.Split(t.In(PST).String(), " ")
  fmt.Fprintln(file, result[0])
  if err != nil {
    fmt.Println(err)
    file.Close()
    return
  }
  for i := 0; i < 35; i++ {
    for j := 0; j < 35; j++ {
      var data = Dist_matrix[i][j]
      s64 := strconv.FormatFloat(data, 'f', -1, 64)
      fmt.Fprintln(file, s64)
    }
  }
  err = file.Close()
  if err != nil {
    fmt.Println(err)
    return
  }
}
