package main

import (
	"github.com/martini-contrib/render"
	"net/http"
	_ "time"
)

type Routes struct {
}

func (routes Routes) landing(renderer render.Render, r *http.Request) {
	data := struct {
		MenuItem string
	} {
		"home",
	}
	renderer.HTML(200, "landing", data)
}

func (routes Routes) discussionIndex(renderer render.Render, r *http.Request) {
	data := struct {
		MenuItem string
	}{
		"talk",
	}
	renderer.HTML(200, "discussions/index", data)
}

func (route Routes) roadmap(renderer render.Render, r *http.Request) {
	data := struct {
		MenuItem string
	}{
		"road",
	}

	renderer.HTML(200, "roadmap", data)
}
