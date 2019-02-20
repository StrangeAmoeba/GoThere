package data_types

type Vertex struct {
  Lat, Long float64
}

/* google api - json struct starts here */

type dir_info struct {
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
  NorthEast LatLong `json:"northeast"`
  SouthWest LatLong `json:"southwest"`
}

type LatLong struct {
  Lat float64  `json:"lat"`
  Long float64 `json:"lng"`
}

type Leg struct {
  Distance      Info           `json:"distance"`
  Duration      Info           `json:"duration"`
  EndAddress    string         `json:"end_address"`
  EndLocation   LatLong        `json:"end_location"`
  StartAddress  string         `json:"start_address"`
  StartLocation LatLong        `json:"start_location"`
  Steps         []Step         `json:"steps"`
  ViaWaypoint   []ViaWaypoint  `json:"via_waypoint"`
}

type Info struct {
  Text float64 `json:"text"`
  Val  float64 `json:"value"`
}

type Step struct {
  Distance         Info           `json:"distance"`
  Duration         Info           `json:"duration"`
  EndLocation      LatLong        `json:"end_location"`
  HtmlInstructions string         `json:"html_instructions"`
  Maneuver         string         `json:"maneuver"`
  Polyline         Point          `json:"polyline"`
  StartLocation    LatLong        `json:"start_location"`
  TravelMode       string         `json:"travel_mode"`
}

type ViaWaypoint struct {
  Location          LatLong  `json:"location"`
  StepIndex         int      `json:"step_index"`
  StepInterpolation float64  `json:"step_interpolation"`
}

type Point struct {
  Pt string `json:"point"`
}

/* google api - json struct ends here */
