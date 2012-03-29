# = require bootstrap
# = require prettify
# = require_tree ./prettify

 $(function() {
    var pusher = new Pusher('c76a8d5a8c95fca62482'),
    channel = pusher.subscribe('trail');
    channel.bind('request',
    function(request) {

        var $request = $('<pre/>', {
            'class': 'prettyprint'
        }).html(JSON.stringify(request, null, 2));

        $request.hide().prependTo('#trail');
        window.prettyPrint();
	    $request.fadeIn('slow');
    });

    $('#example').click(function(e) {
        e.preventDefault();
        $.get('http://httptrail.herokuapp.com/hello');
    });
});
