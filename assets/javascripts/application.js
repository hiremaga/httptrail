#= require bootstrap

$(function() {
  var pusher = new Pusher('c76a8d5a8c95fca62482');
  var channel = pusher.subscribe('trail');
  channel.bind('request', function(request) {
    $('<li/>').html(JSON.stringify(request)).hide().prependTo('#trail').fadeIn('slow');
  });

  $('#example').click(function(e) {
    e.preventDefault();
    $.get('http://httptrail.herokuapp.com/hello');
  });
});
