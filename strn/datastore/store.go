package store

import (
	"github.com/thanyawzinmin/urlstrn/strn/config"
	"gopkg.in/redis.v3"
)

var RedisCon = app.RedisAddress + ":" + app.RedisPort

func CreateDis(uid, url string) {

	// set its address
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     RedisCon,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// test set the string value
	// find the value with key "url"
	err := RedisClient.Set(uid, []byte(url), 0).Err()

	if err != nil {
		panic(err)
	}
}

func GetUrl(uid string) string {
	// set its address
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     RedisCon,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// find the url with the key "uid"
	url, err := RedisClient.Get(uid).Result()

	if err != nil {
		panic(err)
	}

	return string(url)
}
