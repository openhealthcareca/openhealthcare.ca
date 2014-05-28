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
	d := Discussions{}

	server.Get("/", r.landing)

	server.Get("/discuss", d.index)
	server.Get("/discuss/submit", d.create)
	server.Get("/discuss/:id", d.show)
	server.Post("/discuss/store", d.store)
	server.Put("/discuss/:id", d.update)

	server.Get("/roadmap", r.roadmap)

	server.NotFound(func(ren render.Render) {
		ren.HTML(404, "errors/404", nil)
	})

	server.Run()
}
