package main

import (
	"github.com/martini-contrib/render"
	"time"
)

type Routes struct {
}

func (routes Routes) response(r render.Render) {
	
}

func (routes Routes) landing(r render.Render) {
	d := []string{"Where in the world is Carmen San Diego?", "What the hell is derp?"}
	r.HTML(200, "landing", d)
}

func (routes Routes) discussionIndex(r render.Render) {
	data := struct {
		CurrentUrl string
		CurrentDate string 
	} {
		"/talk",
		time.Now().Format("Jan 2, 2006 at 3:03pm (EST)"),
	}
	r.HTML(200, "discussions/index", data)
}
