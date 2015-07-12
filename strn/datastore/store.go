package store

import (
	"github.com/astaxie/goredis"
	"github.com/thanyawzinmin/urlstrn/strn/config"
)

// create goredis instance
var redisClient goredis.Client

func CreateDis(uid, url string) {

	// set its address
	redisClient.Addr = app.RedisAddr

	// test set the string value
	// find the value with key "url"
	redisClient.Set(uid, []byte(url))
	// redisClient.Del(uid)
}

func GetUrl(uid string) string {

	// find the url with the key "uid"
	url, _ := redisClient.Get(uid)
	return string(url)
}
