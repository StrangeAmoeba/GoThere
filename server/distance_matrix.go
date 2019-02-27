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
var Dist_slice_matrix = [][]float64{}
var avg_speed = 15.5 // mps
var updated_matrix = false

// Create_dist_matrix is responsible for assigning weights between each location
// which is important for the application of dijkstas and kruskals algorithm to find the shortest
// route. The matrix created can later be accessed using Dist_matrix which is of the type [][]float64
// We use go routines to create the matrix as quickly as possible.
// WaitGroup is used to find if all go routines have completed their execution.
//
// Input: None
// Output: None
func Create_dist_matrix() {
  check_matrix_file("new")

  // if matrix has been updated this day, return
  if updated_matrix == true {
    Dist_slice_matrix = matToDynMat()
    return
  }

  var wg sync.WaitGroup

  keys := make([]string, 0)
  for k_i := range Locations() {
    keys = append(keys, k_i)
  }
  // Indexing of locations is done in sorted order. So we have to traverse in
  // sorted order for processing user queries.
  sort.Strings(keys)

  for k_i, v_i := range keys {
    // counter for go routines
    wg.Add(1)

    // call go routine
    go func(k_i int, v_i string, keys []string) {
      defer wg.Done() // Decrement the counter when the goroutine completes.
      for k_j, v_j := range keys {
        if updated_matrix == true {
          break
        }
        if k_i != k_j {
          dist_traffic(k_i, v_i, k_j, v_j) // json_parse_dist_traff - assign values to the matrix
        }
        time.Sleep(100 * time.Millisecond)
      }
    }(k_i, v_i, keys)
  }

  // Wait for all go routines to complete
  wg.Wait()

  // fmt.Println("check", Dist_matrix) // debugging
  // fall back to old log incase google api is down or resulted in an error
  if updated_matrix == true {
    check_matrix_file("old")
    Dist_slice_matrix = matToDynMat()
    return
  }

  write_matrix_file()
  Dist_slice_matrix = matToDynMat()
}

// matToDynMat is a helper function which converts the 35*35 matrix (Dist_matrix)
// to a dynamic [][]float64 matrix in accordance with kruskals requirement.
//
// Input: None
// Output: weight[ weight of edge between the two locations ] i.e. float64
func matToDynMat() [][]float64 {
  var mat [][]float64

  for i := 0; i < 35; i++ {
    var row = make([]float64, 35)
    for j := 0; j < 35; j++ {
      row[j] = Dist_matrix[i][j]
    }
    mat = append(mat, row)
  }

  return mat
}

// assign_weight is responsible to normalize the two weights - distance and traffic
// into one weight.
//
// Input: dist[ weight representing distance of route ] i.e. float64,
// traff[ weight representing traffic in route ] i.e. float64
// Output: weight[ weight of edge between the two locations ] i.e. float64
func assign_weight(dist, traff float64) float64 {
  var weight = dist
  weight += (dist * avg_speed)

  // return in km's
  return weight / 1000
}

// check_matrix_file is responsible to check if the current date stamp on file
// is different from current date.(Pacific Time Zone)
// If the current date stamp matches with current date, then update Dist_matrix
// else update updated_matrix accordingly.
//
// Input: f_type [ log to be checked - old or new ] i.e. string
// Output: None
func check_matrix_file(f_type string) {
  var file *os.File
  var err error

  if strings.Compare(f_type, "new") == 0 {
    file, err = os.Open("dist_matrix.log")
  } else {
    file, err = os.Open("dist_matrix.log.old")
  }

  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  // PST date to be verified, IST to process current file.
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
      scanner.Scan()
      var data = scanner.Text()
      if val, err := strconv.ParseFloat(data, 64); err == nil {
        Dist_matrix[i][j] = val
      } else {
        fmt.Println(err)
        os.Exit(3)
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
//
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
