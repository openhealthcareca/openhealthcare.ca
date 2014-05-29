package main

import (
	"github.com/garyburd/redigo/redis"
)

type Persistable interface {
	getRedisConnection() redis.Conn
	persist() bool
}

type RedisPersistable struct{}

func (r RedisPersistable) getConnection() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}

	return c
}
