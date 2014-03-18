SHELL := /bin/bash

all:
	@echo "Openhealthcare installation."
	@ceho " "
	@echo "Install     - Install all dependencies"

install:
	go get github.com/codegangsta/martini
	go get github.com/codegangsta/gin

