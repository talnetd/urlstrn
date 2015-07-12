package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thanyawzinmin/urlstrn/strn/datastore"
	"math/rand"
	"net/http"
)

var keys = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func StoreUrl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	uid := generateId(7)

	// test store
	store.CreateDis(uid, url)
	fmt.Println(url, uid)
	http.Redirect(w, r, "/get/"+uid, http.StatusFound)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	// get request params
	// and assign uid
	vars := mux.Vars(r)
	uid := vars["key"]

	// firl url with
	// key value uid
	rUrl := store.GetUrl(uid)
	http.Redirect(w, r, rUrl, http.StatusFound)
}

func generateId(n int) string {
	gString := make([]rune, n)
	for i := range gString {
		gString[i] = keys[rand.Intn(len(keys))]
	}
	return string(gString)
}
