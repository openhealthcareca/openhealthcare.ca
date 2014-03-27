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
		TestRange []int
	} {
		req,
		"talk",
		filter,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	ren.HTML(200, "discussions/index", data)
}

func (d Discussion) store(ren render.Render, req *http.Request) {

}
