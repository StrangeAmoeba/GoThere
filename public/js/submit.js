$( "#go-form" ).submit(function( event ) {

  // Stop form from submitting normally
  event.preventDefault();

  // Get some values from elements on the page:
  var $form = $( this ),
    term = $('form').serializeArray(),
    url = $form.attr( "action" );

  // Send the data using post
  var posting = $.post( url, { "form_data": term } );
  // Put the results in a result div
  posting.done(function( data ) {
    console.log(data);
    loc_obj = JSON.parse(data);
    console.log(loc_obj);

    // loc_obj consists of info, lat, long of a place respectively, for paths and route_helpers
    initMap(loc_obj);
    // var content = data // debugging
    // $("#result").empty().append(content); // debugging
  });
});
