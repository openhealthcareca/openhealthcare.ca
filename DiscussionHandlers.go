package main

import (
	"github.com/martini-contrib/render"
	"github.com/mholt/binding"
	"net/http"
)

type DiscussionHandlers struct{}

func (d DiscussionHandlers) index(r render.Render, req *http.Request) {
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
		Request     *http.Request
		MenuItem    string
		SubMenuItem string
		TestRange   []int
	}{
		req,
		"discuss",
		filter,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	r.HTML(200, "discussions/index", data)
}

func (d DiscussionHandlers) create(r render.Render, req *http.Request) {
	data := struct {
		Request     *http.Request
		MenuItem    string
		SubMenuItem string
	}{
		req,
		"discuss",
		"submit",
	}

	r.HTML(200, "discussions/create", data)
}

func (d DiscussionHandlers) store(resp http.ResponseWriter, req *http.Request) {
	discussion := new(Discussion)
	errs := binding.Bind(req, discussion)
	if errs.Handle(resp) {
		return
	}

	p := discussion.Persist()

	if p != true {
		return
	}

	http.Redirect(resp, req, "/discussions", http.StatusFound)

	return
}

func (d DiscussionHandlers) show(r render.Render, req *http.Request) {

}

func (d DiscussionHandlers) update(r render.Render, req *http.Request) {

}
