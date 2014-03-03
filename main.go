package main

import (
	"encoding/json"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"github.com/hiremaga/httptrail/sse"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	sseBroker := sse.NewBroker()
	go sseBroker.Start()

	m.Get("/events/", sseBroker.Handler())

	m.Get("/**", func(params martini.Params) (int, string) {
		paramsJSON, err := json.Marshal(params)
		if err != nil {
			return 500, err.Error()
		}

		sseBroker.Messages <- string(paramsJSON)

		return 200, string(paramsJSON)
	})

	m.Run()
}
