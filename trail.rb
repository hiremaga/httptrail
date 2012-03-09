require 'sinatra'

get '/' do
<<-HTML
  <html>
    <head>
      <title>HTTP Trail</title>
      <script src="http://code.jquery.com/jquery-1.7.1.min.js"></script>
      <script src="http://js.pusher.com/1.11/pusher.min.js"></script>
    </head>
    <body>
      <h1>HTTP Trail</h1>
      
      <h2>Try this:</h2>
      <code>
        curl -i <a id='example' href="#">http://httptrail.heroku.com/hello</a>
      </code>

      <ul id="trail"></ul>

      <script>
        $(function() {
          var pusher = new Pusher('c76a8d5a8c95fca62482');
          var channel = pusher.subscribe('trail');
          channel.bind('request', function(request) {
            $('<li/>').html(JSON.stringify(request)).hide().prependTo('#trail').fadeIn('slow');
          });

          $('#example').click(function(e) {
            e.preventDefault();
            $.get('http://httptrail.heroku.com/hello');
          });
        });
      </script>
    </body>
  </html>
HTML
end

get '/favicon.ico' do
  status 404 # Avoids noise from automatic /favicon.ico requests by the browser
end

get '*' do
  require 'cgi'
  require 'pusher'

  Pusher.app_id = '16655'
  Pusher.key    = 'c76a8d5a8c95fca62482'
  Pusher.secret = 'ac24c5833b2197473312'

  data = {received: Time.now, request: request.env}
  Pusher['trail'].trigger('request', data)

  "Request successfully sent to http://httptrail.heroku.com/"
end
