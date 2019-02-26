$( "#go-form" ).submit(function( event ) {

  // Stop form from submitting normally
  event.preventDefault();

  // Get some values from elements on the page:
  console.log("hi");
  console.log($('form').serializeArray())
  var $form = $( this ),
    term = $('form').serializeArray(),
    url = $form.attr( "action" );

  // Send the data using post
  var posting = $.post( url, { "form_data": term } );

  // // Put the results in a div
  // posting.done(function( data ) {
  //   var content = $( data ).find( "#content" );
  //   $( "#result" ).empty().append( content );
  // });
});
