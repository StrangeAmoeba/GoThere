function initMap(loc_obj) {
  // loc_obj consists of info, lat, long of a place respectively.

  var map = new google.maps.Map(document.getElementById('map'), {
    zoom: 11,
    center: new google.maps.LatLng(17.386368, 78.482729),
    mapTypeId: google.maps.MapTypeId.ROADMAP
  });

  var infowindow = new google.maps.InfoWindow({});

  console.log("dedbb", loc_obj, loc_obj["path"].length)

  var marker, i, length = loc_obj["path"].length;

  for (i = 0; i < length; i++) {
    marker = new google.maps.Marker({
      position: new google.maps.LatLng(loc_obj["path"][i][1], loc_obj["path"][i][2]),
      map: map
    });

    google.maps.event.addListener(marker, 'click', (function (marker, i) {
      return function () {
        infowindow.setContent(loc_obj["path"][i][0]);
        infowindow.open(map, marker);
      }
    })(marker, i));
  }
}
