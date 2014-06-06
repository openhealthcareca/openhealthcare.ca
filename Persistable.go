package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type Persistable interface {
	getConnection() redis.Conn
	Persist() bool
}

type RedisPersistable struct{}

func (r RedisPersistable) getConnection() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	c.Do("SELECT", "1")

	if err != nil {
		panic(err)
	}

	return c
}

func (r RedisPersistable) getNewId(uid string) int {
	c := r.getConnection()

	defer c.Close()

	_, err := c.Do("INCR", uid)

	if err != nil {
		log.Fatal(err)
	}

	id, err := redis.Int(c.Do("GET", uid))

	if err != nil {
		log.Fatal(err)
	}
	return id
}
