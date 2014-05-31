package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
)

type Discussions struct {
	RedisPersistable
	all map[int]Discussion
}

func (d *Discussions) Retrieve(id int) *Discussion {

	c := d.getConnection()

	if id <= 0 {
		log.Print("Invalid discussion id")
	}

	existingDiscussion, err := redis.Values(c.Do("HGETALL", "discussions:"+strconv.Itoa(id)))

	if err != nil {
		log.Fatal(err)
	}

	discussion := new(Discussion)
	if err := redis.ScanStruct(existingDiscussion, &discussion); err != nil {
		log.Fatal(err)
	}

	return discussion
}
