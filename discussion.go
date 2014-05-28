package main

import (
	"github.com/martini-contrib/render"
	"net/http"
	 _ "time"
)

type Discussions struct {}

func (d Discussions) index(ren render.Render, req *http.Request) {
	// If the request is to 
	filter := req.URL.Query().Get("filter")
	if "all" == filter {
		// Grab all the posts
	} else if "trending" == filter {
		// Grab only trending posts
	} else if "new" == filter {
		// Grab only new posts
	}

	data := struct {
		Request *http.Request
		MenuItem string
		SubMenuItem string
		TestRange []int
	} {
		req,
		"discuss",
		filter,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	ren.HTML(200, "discussions/index", data)
}

func (d Discussions) create(ren render.Render, req *http.Request) {
	data := struct {
		Request *http.Request
		MenuItem string
		SubMenuItem string
	} {
		req,
		"discuss",
		"submit",
	}
	
	ren.HTML(200, "discussions/create", data)
}

func (d Discussions) store(ren render.Render, req *http.Request) {

}

func (d Discussions) show(ren render.Render, req *http.Request) {
	
}

func (d Discussions) update(ren render.Render, req *http.Request) {
	
}

