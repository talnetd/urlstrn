package api

import (
	"fmt"
	"math/rand"
	"net/http"
)

var keys = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func GetUrl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	uid := generateId(7)

	fmt.Println(url, uid)
}

func generateId(n int) string {
	gString := make([]rune, n)
	for i := range gString {
		gString[i] = keys[rand.Intn(len(keys))]
	}
	return string(gString)
}
