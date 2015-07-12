package home

import (
	"github.com/gorilla/mux"
	"github.com/thanyawzinmin/urlstrn/strn/config"
	"github.com/thanyawzinmin/urlstrn/strn/datastore"
	"html/template"
	"log"
	"net/http"
	"os"
)

type HomeVariables struct {
	Title    string
	FullLink string
	OrgLink  string
}

func MakeHome(w http.ResponseWriter, r *http.Request) {
	type HomeVariables struct {
		Title string
	}

	p := HomeVariables{
		Title: "Strn: as simple as your hands",
	}

	app.Templates = template.Must(template.ParseFiles("strn/templates/home/home.html", app.MainLayout))
	app.Templates.ExecuteTemplate(w, "base", p)
}

func ShowStrn(w http.ResponseWriter, r *http.Request) {
	type GetVariables struct {
		Title     string
		OrgLink   string
		FullLink  string
		RedisPort string
	}
	vars := mux.Vars(r)
	uid := vars["key"]

	p := GetVariables{
		Title:     "Have been Shortened",
		FullLink:  r.Host + "/" + uid,
		OrgLink:   store.GetUrl(uid),
		RedisPort: os.Getenv("REDIS_PORT_6379_TCP_ADDR"),
	}

	log.Println(p)
	log.Println(os.Getenv("REDIS_PORT_6379_TCP_ADDR"))

	app.Templates = template.Must(template.ParseFiles("strn/templates/home/get.html", app.MainLayout))
	app.Templates.ExecuteTemplate(w, "base", p)
}
