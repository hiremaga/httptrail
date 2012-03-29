require 'sinatra'

get '/' do
  erb :index
end

get '/favicon.ico' do
  status 404 # Avoids noise from automatic /favicon.ico requests by the browser
end

get '*' do
  require 'cgi'
  require 'pusher'

  unless ENV['RACK_ENV'] == 'production'
    Pusher.app_id = '17726'
    Pusher.key    = '02362921b5f72e0fa84a'
    Pusher.secret = '919a494088a3a998c81b'
  end

  data = {received: Time.now, request: request.env}
  Pusher['trail'].trigger('request', data)

  "Request successfully sent to http://httptrail.heroku.com/"
end
