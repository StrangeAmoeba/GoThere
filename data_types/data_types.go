package data_types

type Vertex struct {
  Lat, Long float64
}

type Weight_tuple struct {
  Speed, Distance float64
}

/* google api - json struct starts here */

type Dir_info struct {
  Waypoints []Waypoint `json:"geocoded_waypoints"`
  Routes    []Route    `json:"routes"`
  Status    string     `json:"status"`
}

type Waypoint struct {
  Stat     string   `json:"geocoder_status"`
  id       string   `json:"place_id"`
  types    []string `json:"types"`
}

type Route struct {
  Bound       []Corner    `json:"bounds"`
  Cp          string      `json:"copyrights"`
  Legs        []Leg       `json:"legs"`
  Ov_polyline Point  `json:"overview_polyline"` // overview polyline
  Summary     string      `json:"summary"`
  Warnings    []string    `json:"warnings"`
  Order       []int        `json:"waypoint_order"` // waypoint order
}

type Corner struct {
  NorthEast Lat_long `json:"northeast"`
  SouthWest Lat_long `json:"southwest"`
}

type Lat_long struct {
  Lat float64  `json:"lat"`
  Long float64 `json:"lng"`
}

type Leg struct {
  Distance      Info           `json:"distance"`
  Duration      Info           `json:"duration"`
  End_address    string         `json:"end_address"`
  End_location   Lat_long        `json:"end_location"`
  Start_address  string         `json:"start_address"`
  Start_location Lat_long        `json:"start_location"`
  Steps         []Step         `json:"steps"`
  Via_waypoint   []Via_waypoint  `json:"via_waypoint"`
}

type Info struct {
  Text float64 `json:"text"`
  Val  float64 `json:"value"`
}

type Step struct {
  Distance         Info           `json:"distance"`
  Duration         Info           `json:"duration"`
  End_eocation      Lat_long        `json:"end_location"`
  Html_instructions string         `json:"html_instructions"`
  Maneuver         string         `json:"maneuver"`
  Polyline         Point          `json:"polyline"`
  Start_location    Lat_long        `json:"start_location"`
  Travel_mode       string         `json:"travel_mode"`
}

type Via_waypoint struct {
  Location          Lat_long  `json:"location"`
  Step_index         int      `json:"step_index"`
  Step_interpolation float64  `json:"step_interpolation"`
}

type Point struct {
  Pt string `json:"point"`
}

/* google api - json struct ends here */

type Graph_edge struct {
  Src,Dst int
  Weight float64
}