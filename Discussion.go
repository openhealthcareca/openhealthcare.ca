package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/mholt/binding"
	"log"
	"strconv"
	"time"
)

type Discussion struct {
	RedisPersistable
	Id      string
	Title   string
	URL     string
	Text    string
	Created time.Time
	Updated time.Time

	key string
	uid string
}

func NewDiscussion() *Discussion {
	return &Discussion{
		key: "discussions",
		uid: "discussions:uid",
	}
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
		d.Id = strconv.Itoa(d.RedisPersistable.getNewId(d.uid))
		// Add a new discussion to the list of them
		c.Do("SADD", "discussions", d.Id)
	}

	if _, err := c.Do("HMSET", redis.Args{}.Add("discussions:"+d.Id).AddFlat(d)...); err != nil {
		log.Fatal(err)
	}

	return true
}
