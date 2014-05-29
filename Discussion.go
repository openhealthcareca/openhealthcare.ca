package main

import (
	"github.com/mholt/binding"
	_ "time"
)

type Discussion struct {
	RedisPersistable
	Id    int
	Title string
	URL   string
	Text  string
}

func (d *Discussion) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&d.Id:    "id",
		&d.Title: "title",
		&d.URL:   "URL",
		&d.Text:  "text",
	}
}

func (d *Discussion) persist() bool {

	c := d.getConnection()
	defer c.Close()

	if d.Id != 0 {
		d.Id = d.getNewId()
	}

	c.Do("SET", "discussions:"+string(d.Id)+":title", d.Title)

	return true
}

func (d *Discussion) getNewId() int {
	c := d.getConnection()

	defer c.Close()

	id, err := c.Do("INCR", "discussions:uid")

	if err != nil {
		panic(err)
	}

	return id.(int)
}
