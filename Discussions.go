package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
)

type Discussions struct {
	RedisPersistable
	All map[int]Discussion
}

func (ds Discussions) retrieveAll() []*Discussion {
	c := ds.getConnection()

	defer c.Close()

	allDiscussions, err := redis.Values(c.Do("SMEMBERS", "discussions"))

	if err != nil {
		log.Fatal(err)
	}

	discussions := make([]*Discussion, 1)

	dss := []int{}

	if err := redis.ScanSlice(allDiscussions, &dss); err != nil {
		log.Fatal(err)
	}

	for _, id := range dss {

		discussions = append(discussions, ds.retrieveById(id))
	}

	log.Print(discussions)

	return discussions
}

func (ds Discussions) retrieveById(id int) *Discussion {

	c := ds.getConnection()

	defer c.Close()

	if id <= 0 {
		log.Fatal("Cannot retrieve ID" + strconv.Itoa(id))
	}

	item, err := redis.Values(c.Do("HGETALL", "discussions:"+strconv.Itoa(id)))

	if err != nil {
		log.Fatal(err)
	}

	d := new(Discussion)

	if err := redis.ScanStruct(item, d); err != nil {
		log.Fatal(err)
	}

	return d
}
