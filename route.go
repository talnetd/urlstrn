package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/thanyawzinmin/urlstrn/strn/handlers/api"
	"github.com/thanyawzinmin/urlstrn/strn/handlers/home"
	"log"
	"net/http"
)

func main() {
	// load environment files
	godotenv.Load()

	// create a new mux router
	request := mux.NewRouter()

	// asign a handler function to root "/"
	request.HandleFunc("/", home.MakeHome)

	// main entry home view
	http.Handle("/", request)

	// end point for api
	// * for all clients
	// mobile* platform* web*
	request.HandleFunc("/api/url", api.StoreUrl).Methods("GET", "POST")

	// [web] after saving the strn link
	// show it back to user through a page
	request.HandleFunc("/get/{key}", home.ShowStrn).Methods("GET")
	request.HandleFunc("/{key}", api.Redirect).Methods("GET")

	// to serve staic contents
	// *css *js *html may be
	assetServer := http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	http.Handle("/public/", assetServer)

	// dependency management done by bower
	bowerServer := http.StripPrefix("/bower_components/", http.FileServer(http.Dir("bower_components")))
	http.Handle("/bower_components/", bowerServer)

	// Start listening on port :3000
	log.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
