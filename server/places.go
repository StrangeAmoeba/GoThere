//
// Description: Stores the latitute and longitude translations of the defined places.
//

package server

type Vertex struct {
  Lat, Long float64
}

var Locations = func() map[string] Vertex { return map[string] Vertex {
    "ahmedguda": Vertex{
      17.517813, 78.599766 },
    "allwyncolony": Vertex{
      17.504999, 78.413737 },
    "ankireddipalli": Vertex{
      17.564021, 78.684769 },
    "ameenpur": Vertex{
      17.520516, 78.319846 },
    "ameerpet": Vertex{
      17.437599, 78.448241 },
    "bachupally": Vertex{
      17.549203, 78.385600 },
    "begumpet": Vertex{
      17.444339, 78.469836 },
    "bolarum": Vertex{
      17.531032, 78.514498 },
    "borabanda": Vertex{
      17.457145, 78.412862 },
    "bowenpally": Vertex{
      17.473768, 78.487278 },
    "bhel": Vertex{
      17.491567, 78.292227 },
    "brundavancolony": Vertex{
      17.527584, 78.492173 },
    "chandanagar": Vertex{
      17.494780, 78.325199 },
    "dammaiguda": Vertex{
      17.499497, 78.592886 },
    "fatehnagar": Vertex{
      17.457844, 78.452139 },
    "gachibowli": Vertex{
      17.442694, 78.354830 },
    "gautamnagar": Vertex{
      17.462244, 78.457899 },
    "gopanapalli": Vertex{
      17.457720, 78.313973 },
    "hafeezpet": Vertex{
      17.483392, 78.363973 },
    "hiteccity": Vertex{
      17.448928, 78.379135 },
    "hydernagar": Vertex{
      17.501590, 78.380543 },
    "ismailkhanpet": Vertex{
      17.628537, 78.158098 },
    "jubileehills": Vertex{
      17.431963, 78.407018 },
    "kandi": Vertex{
      17.584131, 78.107609 },
    "kompally": Vertex{
      17.542415, 78.483240 },
    "kondapur": Vertex{
      17.561204, 78.009920 },
    "kothaguda": Vertex{
      17.463256, 78.376216 },
    "kukatpally": Vertex{
      17.484393, 78.390273 },
    "lingampally": Vertex{
      17.489642, 78.316544 },
    "madhapur": Vertex{
      17.443359, 78.393705 },
    "miyapur": Vertex{
      17.509155, 78.352031 },
    "motinagar": Vertex{
      17.449186, 78.418215 },
    "nallagandla": Vertex{
      17.467881, 78.309822 },
    "nizampet": Vertex{
      17.517239, 78.380958 },
    "patancheru": Vertex{
      17.528927, 78.266828 },
    "patelnagar": Vertex{
      17.467252, 78.421853 },
    "peddakanjerla": Vertex{
      17.600183, 78.244732 },
    "pragathinagar": Vertex{
      17.521026, 78.396227 },
    "dullapally": Vertex{
      17.553054, 78.463240 },
    "sainagar": Vertex{
      17.547152, 78.497131 },
    "sanathnagar": Vertex{
      17.456273, 78.444159 },
    "sangareddy": Vertex{
      17.620911, 78.081974 },
    "shamshiguda": Vertex{
      17.506510, 78.403681 },
    "sivanagar": Vertex{
      17.610880, 78.273745 },
    "suraram": Vertex{
      17.545988, 78.434571 },
    "tellapur": Vertex{
      17.462154, 78.286401 },
    "turkapalli": Vertex{
      17.547217, 78.513067 },
    "whispervalley": Vertex{
      17.540640, 78.371781 },
    "yeddumailaram": Vertex{
      17.500511, 78.135255 },
    "yousufguda": Vertex{
      17.433208, 78.429271 },
  }
}
