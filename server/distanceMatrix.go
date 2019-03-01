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

var DistMatrix = [35][35]float64{}
var DistSliceMatrix = [][]float64{}
var updatedMatrix = false

// AvgSpeed can be manually changed by user.
// Used value based on reports from 2018.
var AvgSpeed = 15.5 // mps

// CreateDistMatrix is responsible for assigning weights between each location
// which is important for the application of dijkstas and kruskals algorithm to find the shortest
// route. The matrix created can later be accessed using DistMatrix which is of the type [][]float64
// We use go routines to create the matrix as quickly as possible.
// WaitGroup is used to find if all go routines have completed their execution.
//
//  Input: None
//  Output: None
func CreateDistMatrix() {
  CheckMatrixFile()

  // if matrix has been updated this day, return
  if updatedMatrix == true {
    DistSliceMatrix = MatToDynMat()
    return
  }

  var wg sync.WaitGroup

  keys := make([]string, 0)
  for kI := range Locations() {
    keys = append(keys, kI)
  }
  // Indexing of locations is done in sorted order. So we have to traverse in
  // sorted order for processing user queries.
  sort.Strings(keys)

  for kI, vI := range keys {
    // counter for go routines
    wg.Add(1)

    // call go routine
    go func(kI int, vI string, keys []string) {
      defer wg.Done() // Decrement the counter when the goroutine completes.
      for kJ, vJ := range keys {
        if updatedMatrix == true {
          break
        }
        if kI != kJ {
          DistTraffic(kI, vI, kJ, vJ) // jsonParseDistTraff - assign values to the matrix
        }
        time.Sleep(200 * time.Millisecond)
      }
    }(kI, vI, keys)
  }

  // Wait for all go routines to complete
  wg.Wait()

  // fmt.Println("check", DistMatrix) // debugging
  // fall back to log incase google api is down or resulted in an error
  if updatedMatrix == true {
    fmt.Println("OVER QUERY LIMIT")
    CheckMatrixFile()
    DistSliceMatrix = MatToDynMat()
    return
  }

  WriteMatrixFile()
  DistSliceMatrix = MatToDynMat()
}

// MatToDynMat is a helper function which converts the 35*35 matrix (DistMatrix)
// to a dynamic [][]float64 matrix in accordance with kruskals requirement.
//
//  Input: None
//  Output: weight[ weight of edge between the two locations ] i.e. float64
func MatToDynMat() [][]float64 {
  var mat [][]float64

  for i := 0; i < 35; i++ {
    var row = make([]float64, 35)
    for j := 0; j < 35; j++ {
      row[j] = DistMatrix[i][j]
    }
    mat = append(mat, row)
  }

  return mat
}

// AssignWeight is responsible to normalize the two weights - distance and traffic
// into one weight.
//
//  Input: dist[ weight representing distance of route ] i.e. float64,
//         traff[ weight representing traffic in route ] i.e. float64
//  Output: weight[ weight of edge between the two locations ] i.e. float64
func AssignWeight(dist, traff float64) float64 {
  var weight = dist
  weight += (dist * AvgSpeed)

  // return in km's
  return weight / 1000
}

// CheckMatrixFile is responsible to check if the current date stamp on file
// is different from current date.(Pacific Time Zone)
// If the current date stamp matches with current date, then update DistMatrix
// else update updatedMatrix accordingly.
//
//  Input: None
//  Output: None
func CheckMatrixFile() {
  var file *os.File
  var err error

  file, err = os.Open("distMatrix.log")

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
  updatedMatrix = true
  for i := 0; i < 35; i++ {
    for j := 0; j < 35; j++ {
      scanner.Scan()
      var data = scanner.Text()
      if val, err := strconv.ParseFloat(data, 64); err == nil {
        DistMatrix[i][j] = val
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

// WriteMatrixFile is responsible to store the generated weight matrix into a file for accessing
// the matrix on the same day later again. Saves api requests made to google api.
// Only writes if the current date stamp on file is different from current date.(Pacific Time Zone)
//
//  Input: None
//  Output: None
func WriteMatrixFile() {
  file, err := os.Create("distMatrix.log")
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
      var data = DistMatrix[i][j]
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
