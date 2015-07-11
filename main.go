package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thanyawzinmin/urlstrn/strn/redis"
	"log"
	"net/http"
	"time"
)

func main() {
	// create a new mux router
	request := mux.NewRouter()

	// asign a handler function to root "/"
	request.HandleFunc("/", RootHandler)

	// dummy test for routing
	http.Handle("/", request)

	// to serve staic contents
	// *css *js *html may be
	assetServer := http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	http.Handle("/public/", assetServer)

	// Start listening on port :3000
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// get the current time Format used RFC1123
	// >> http://www.csgnetwork.com/timerfc1123calc.html
	tnow := time.Now().Format(time.RFC1123)

	// write out the return
	w.Write([]byte("The request sent at: " + tnow))

	store.CreateStance()
}
