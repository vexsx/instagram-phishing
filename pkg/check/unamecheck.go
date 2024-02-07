package check

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Username(username string) bool {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.instagram.com/"+username, nil)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return false
	}
	bodyStr := string(body)

	switch resp.StatusCode {
	case 404:
		fmt.Printf("\033[32;1m[UNAVAILABLE] https://www.instagram.com/%s\033[0m\n", username)
		return false
	case 200:
		if strings.Contains(bodyStr, "<title>Instagram</title>") {
			fmt.Printf("\033[31;1m[UNAVAILABLE] https://www.instagram.com/%s\u001B[0m\n", username)
			// dialog.Alert("username is incorrect !!!")
			return false
		} else {
			fmt.Printf("\u001B[32;1m[AVAILABLE] https://www.instagram.com/%s\033[0m\n", username)
			return true
		}
	default:
		fmt.Println("[ERROR] Unexpected status code:", resp.StatusCode)
		return false
	}
}
