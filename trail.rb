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

  Pusher.app_id = '16655'
  Pusher.key    = 'c76a8d5a8c95fca62482'
  Pusher.secret = 'ac24c5833b2197473312'

  data = {received: Time.now, request: request.env}
  Pusher['trail'].trigger('request', data)

  "Request successfully sent to http://httptrail.heroku.com/"
end
