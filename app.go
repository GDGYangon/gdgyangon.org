package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "main",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "Hello World", "Hello World Content")
	})

	m.Get("/events", func(r render.Render) {
		r.HTML(200, "GDG( Yangon ) MeetUps", "Nothing yet")
	})

	m.Use(martini.Static("assets"))
	http.Handle("/", m)
}
