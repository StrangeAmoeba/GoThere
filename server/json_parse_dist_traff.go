package server

// dist_traffic uses the google_directions_api and with the help of
// defined struct to parse json, we obtain distance and traffic weights
// calls assign_weight() to normalize the two parameters into one weight.
// Input: loc1[ origin of the route ] i.e. string, loc2[ origin of the route ] i.e. string
// Output: weight[ weight of edge between the two locations ] i.e. float64
func dist_traffic(loc1, loc2 string) float64 {
  var url = constructURL(Locations()[v_i], Locations()[v_j]) // external_api
  var content = getResponse(url)
  var directions dt.Dir_info
  json.Unmarshal([]byte(content), &directions)
  fmt.Println("debug3", directions.Routes[0].Legs[0].Distance.Val)
  var dist = directions.Routes[0].Legs[0].Distance.Val
  var traff = directions.Routes[0].Legs[0].Duration_traffic.Val
  var weight = assign_weight(dist, traff) // distance_matrix
  return weight
  // getRespFile() - for debugging only
}

// getRespFile is a helper function
// useful for debugging. Can access json file to obtain the parsed data.
// uses (server/example-route.json).
// Input: None
// Output: None
func getRespFile() {
  jsonFile, err := os.Open("server/example-route.json")
  // if we os.Open returns an error then handle it
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Successfully Opened json file")
  // defer the closing of our jsonFile so that we can parse it later on
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  var directions dt.Dir_info
  json.Unmarshal([]byte(byteValue), &directions)
}
