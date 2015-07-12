package store

import (
	"fmt"
	"github.com/astaxie/goredis"
	"github.com/thanyawzinmin/urlstrn/strn/config"
)

func StoreDis() {

	// create goredis instance
	var redisClient goredis.Client

	// set its address
	redisClient.Addr = app.RedisAddr

	// test set the string value
	// find the value with key "url"
	redisClient.Set("url", []byte("http://tinaunglinn.com"))
	query, _ := redisClient.Get("url")

	// print out the result for debug purpose
	// delete the record back
	fmt.Println(string(query))
}
