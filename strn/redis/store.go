package store

import (
	"fmt"
	"github.com/astaxie/goredis"
)

// redis instance addr
const redis = "127.0.0.1:6379"

func CreateStance() {

	// create goredis instance
	var redisClient goredis.Client

	// set its address
	redisClient.Addr = redis

	// test set the string value
	redisClient.Set("url", []byte("http://tinaunglinn.com"))
	// find the value with key "url"
	query, _ := redisClient.Get("url")
	// print out the result for debug purpose
	fmt.Println(string(query))
	// delete the record back
	redisClient.Del("a")
}
