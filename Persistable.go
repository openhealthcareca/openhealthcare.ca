package main

import (
	"github.com/garyburd/redigo/redis"
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
