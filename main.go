package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thanyawzinmin/urlstrn/strn/handlers/api"
	"github.com/thanyawzinmin/urlstrn/strn/handlers/home"
	"log"
	"net/http"
)

func main() {
	// create a new mux router
	request := mux.NewRouter()

	// asign a handler function to root "/"
	request.HandleFunc("/", home.MakeHome)

	// dummy test for routing
	http.Handle("/", request)

	// end point for api
	// * for all clients
	// mobile* platform* web*
	request.HandleFunc("/api/url", api.GetUrl)

	// to serve staic contents
	// *css *js *html may be
	assetServer := http.StripPrefix("/public/", http.FileServer(http.Dir("public")))
	http.Handle("/public/", assetServer)

	// dependency management done by bower
	bowerServer := http.StripPrefix("/bower_components/", http.FileServer(http.Dir("bower_components")))
	http.Handle("/bower_components/", bowerServer)

	// Start listening on port :3000
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
