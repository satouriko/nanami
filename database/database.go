package database

import (
	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

var c redis.Conn

func Init() (err error) {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Get(section string, key string) (values []interface{})  {
	values, err := redis.Values(c.Do("HGETALL", section + "-" + key))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Set(section string, key string, value interface{})  {
	if _, err := c.Do("HMSET", redis.Args{section + "-" + key}.AddFlat(value)...); err != nil {
		log.Fatal(err)
	}
}

func ScanStruct(src []interface{}, dest interface{}) error {
	err := redis.ScanStruct(src, dest)
	return err
}