package home

import (
	"github.com/thanyawzinmin/urlstrn/strn/config"
	"html/template"
	"net/http"
)

func MakeHome(w http.ResponseWriter, r *http.Request) {
	type PageVariables struct {
		Title string
	}

	p := PageVariables{
		Title: "Strn: as simple as your hands",
	}

	app.Templates = template.Must(template.ParseFiles("strn/templates/home/home.html", app.MainLayout))
	app.Templates.ExecuteTemplate(w, "base", p)
}
