package save

import (
	"fmt"
	"os"
)

func SaveCredentials(username, password string) {
	// Open the file in append mode
	staticFolderPath := "./static"
	filePath := staticFolderPath + "/credentials.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
