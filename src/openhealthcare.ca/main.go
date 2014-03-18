package main

import (
	"github.com/codegangsta/martini"
)

func handleRoot() string {

	return "<h1>Openhealthcare.ca</h1>"

}

func main() {
	server := martini.Classic()
	server.Get("/", handleRoot)
	server.Run()
}
