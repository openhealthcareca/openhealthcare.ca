package main

import (
	"github.com/martini-contrib/render"
	"github.com/mholt/binding"
	"net/http"
)

type DiscussionHandlers struct{}

func (d DiscussionHandlers) index(r render.Render, req *http.Request) {

	ds := new(Discussions)

	data := struct {
		Request     *http.Request
		MenuItem    string
		SubMenuItem string
		Discussions []*Discussion
	}{
		req,
		"discuss",
		"all",
		ds.retrieveAll(),
	}

	r.HTML(200, "discussions/index", data)
}

func (d DiscussionHandlers) trending(r render.Render, req *http.Request) {
	data := struct {
		Request *http.Request
	}{
		req,
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
