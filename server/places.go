//
// Description: Stores the latitute and longitude translations of the defined places.
//

package server

import dt "concurrency-9/data_types"

// Locations returns a map[string] of type dt.Vertex.
// Other packages can access these places using Locations()
// Input: none
// Output: A map[string] dt.Vertex
var Locations = func() map[string]dt.Vertex {
  return map[string]dt.Vertex{
    "ahmedguda": dt.Vertex{
      17.517813, 78.599766, 0},
    "allwyncolony": dt.Vertex{
      17.504999, 78.413737, 1},
    "ameenpur": dt.Vertex{
      17.520516, 78.319846, 2},
    "ameerpet": dt.Vertex{
      17.437599, 78.448241, 3},
    "bachupally": dt.Vertex{
      17.549203, 78.385600, 4},
    "begumpet": dt.Vertex{
      17.444339, 78.469836, 5},
    "bhel": dt.Vertex{
      17.491567, 78.292227, 6},
    "borabanda": dt.Vertex{
      17.457145, 78.412862, 7},
    "bowenpally": dt.Vertex{
      17.473768, 78.487278, 8},
    "brundavancolony": dt.Vertex{
      17.527584, 78.492173, 9},
    "chandanagar": dt.Vertex{
      17.494780, 78.325199, 10},
    "fatehnagar": dt.Vertex{
      17.457844, 78.452139, 11},
    "gachibowli": dt.Vertex{
      17.442694, 78.354830, 12},
    "gautamnagar": dt.Vertex{
      17.462244, 78.457899, 13},
    "hafeezpet": dt.Vertex{
      17.483392, 78.363973, 14},
    "hiteccity": dt.Vertex{
      17.448928, 78.379135, 15},
    "hydernagar": dt.Vertex{
      17.501590, 78.380543, 16},
    "jubileehills": dt.Vertex{
      17.431963, 78.407018, 17},
    "kandi": dt.Vertex{
      17.584131, 78.107609, 18},
    "kompally": dt.Vertex{
      17.542415, 78.483240, 19},
    "kondapur": dt.Vertex{
      17.561204, 78.009920, 20},
    "kukatpally": dt.Vertex{
      17.484393, 78.390273, 21},
    "lingampally": dt.Vertex{
      17.489642, 78.316544, 22},
    "madhapur": dt.Vertex{
      17.443359, 78.393705, 23},
    "miyapur": dt.Vertex{
      17.509155, 78.352031, 24},
    "nallagandla": dt.Vertex{
      17.467881, 78.309822, 25},
    "nizampet": dt.Vertex{
      17.517239, 78.380958, 26},
    "patancheru": dt.Vertex{
      17.528927, 78.266828, 27},
    "patelnagar": dt.Vertex{
      17.467252, 78.421853, 28},
    "sainagar": dt.Vertex{
      17.547152, 78.497131, 29},
    "sanathnagar": dt.Vertex{
      17.456273, 78.444159, 30},
    "sangareddy": dt.Vertex{
      17.620911, 78.081974, 31},
    "shamshiguda": dt.Vertex{
      17.506510, 78.403681, 32},
    "sivanagar": dt.Vertex{
      17.610880, 78.273745, 33},
    "yeddumailaram": dt.Vertex{
      17.500511, 78.135255, 34},
  }
}
