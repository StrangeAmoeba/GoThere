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
    obj = JSON.parse(data);
    console.log("debug", obj);
    var content = data
    $("#result").empty().append(content);
  });
});
