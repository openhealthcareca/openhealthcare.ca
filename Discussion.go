package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/mholt/binding"
	"log"
	"strconv"
	_ "time"
)

type Discussion struct {
	RedisPersistable
	Id    string
	Title string
	URL   string
	Text  string
}

func (d *Discussion) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&d.Id:    "id",
		&d.Title: "title",
		&d.URL:   "url",
		&d.Text:  "text",
	}
}

func (d *Discussion) Persist() bool {

	c := d.getConnection()
	defer c.Close()

	if d.Id == "" {
		d.Id = strconv.Itoa(d.getNewId())
		// Add a new discussion to the list of them
		c.Do("SADD", "discussions", d.Id)
	}

	if _, err := c.Do("HMSET", redis.Args{}.Add("discussions:"+d.Id).AddFlat(d)...); err != nil {
		log.Fatal(err)
	}

	return true
}

func (d *Discussion) getNewId() int {
	c := d.getConnection()

	defer c.Close()

	_, err := c.Do("INCR", "discussions:uid")

	if err != nil {
		log.Fatal(err)
	}

	id, err := redis.Int(c.Do("GET", "discussions:uid"))

	if err != nil {
		log.Fatal(err)
	}

	return id

}
