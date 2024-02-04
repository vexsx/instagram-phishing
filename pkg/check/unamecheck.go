package check

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Username(username string) bool {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.instagram.com/"+username+"/", nil)
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

	switch resp.StatusCode {
	case 404:
		fmt.Printf("\033[32;1m[AVAILABLE] https://www.instagram.com/%s\033[0m\n", username)
		return false
	case 200:
		if strings.Contains(string(body), "Login • Instagram") {
			fmt.Println("[!] Failed to check username. Try again later", username)
			return true
		} else {
			fmt.Printf("\033[31;1m[UNAVAILABLE] https://www.instagram.com/%s\033[0m\n", username)
			return false
		}
	default:
		fmt.Println("[ERROR] Unexpected status code:", resp.StatusCode)
		return false
	}
}
