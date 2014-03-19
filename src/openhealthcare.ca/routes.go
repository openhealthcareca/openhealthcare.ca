package main

import (
	"github.com/codegangsta/martini-contrib/render"
)

type Routes struct {}


func (routes Routes) landing(r render.Render) {
	d := []string{"Where in the world is Carmen San Diego?", "What the hell is derp?"}
	r.HTML(200, "landing", d)
}
