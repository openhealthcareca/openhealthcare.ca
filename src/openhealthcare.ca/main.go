package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
)

func main() {
	server := martini.Classic()
	server.Use(gzip.All())

	server.Use(render.Renderer(render.Options{
		Directory: "views",
		Layout: "layout",
	}))
	
	r := Routes{}

	server.Get("/", r.landing)
	server.Get("/talk", r.discussionIndex)
	server.Get("/roadmap", r.roadmap)

	server.Run()
}
