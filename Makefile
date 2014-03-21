SHELL := /bin/bash
GOPATH := $(shell pwd)

all:
	@echo "Openhealthcare installation."
	@ceho " "
	@echo "Install     - Install all dependencies"

install:
	cd $(GOPATH)/src/openhealthcare.ca && go get

