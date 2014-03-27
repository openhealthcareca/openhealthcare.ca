package main

import (
	"github.com/martini-contrib/render"
	"net/http"
	 _ "time"
)

type Discussion struct {}

func (d Discussion) index(ren render.Render, req *http.Request) {
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
		ActiveFilter string
	} {
		req,
		"talk",
		filter,
	}

	ren.HTML(200, "discussions/index", data)
}

func (d Discussion) store(ren render.Render, req *http.Request) {

}
