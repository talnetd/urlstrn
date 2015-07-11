package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	// create a new mux router
	request := mux.NewRouter()

	// asign a handler function to root "/"
	request.HandleFunc("/", RootHandler)

	http.Handle("/", request)

	// Start listening on port :3000
	log.Println("Listening...")
	http.ListenAndServe(":3000", request)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// get the current time Format used RFC1123
	// >> http://www.csgnetwork.com/timerfc1123calc.html
	tnow := time.Now().Format(time.RFC1123)

	// write out the return
	w.Write([]byte("The request sent at: " + tnow))
}
