package main

import (
	"github.com/go-martini/martini"
	"net/http"
)

func init() {
	m := martini.Classic()
	m.Get("/events", func() string {
		return "Events will be here"
	})
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Use(martini.Static("assets"))
	http.Handle("/", m)
}
