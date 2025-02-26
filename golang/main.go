package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getJWTToken() (string, error) {
	// Try to get the token from the environment variable
	token := os.Getenv("auth-jabbercracky")
	if token != "" {
		return strings.TrimSpace(token), nil
	}

	// Try to get the token from the file ~/.jabbercracky
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	tokenFilePath := filepath.Join(homeDir, ".jabbercracky")
	tokenBytes, err := ioutil.ReadFile(tokenFilePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(tokenBytes)), nil
}

func listHashLists() {
	// Get the JWT token
	token, err := getJWTToken()
	if err != nil {
		fmt.Println("Error getting JWT token:", err)
		return
	}

	// Create the request
	req, err := http.NewRequest("GET", "https://jabbercracky.com/api/game/hashlist", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching hash lists:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}

func downloadHashList(id string) {
	// Get the JWT token
	token, err := getJWTToken()
	if err != nil {
		fmt.Println("Error getting JWT token:", err)
		return
	}

	// Create the request
	url := fmt.Sprintf("https://jabbercracky.com/api/game/hashlist/%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching hash list:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}

func submitGameData(id string, filePath string) {
	// Get the JWT token
	token, err := getJWTToken()
	if err != nil {
		fmt.Println("Error getting JWT token:", err)
		return
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create the request
	url := fmt.Sprintf("https://jabbercracky.com/api/game/submit/%s", id)
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error submitting game data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
	submitCmd := flag.NewFlagSet("submit", flag.ExitOnError)

	downloadID := downloadCmd.String("id", "", "Hash List ID")
	submitID := submitCmd.String("id", "", "Hash List ID")
	submitFile := submitCmd.String("file", "", "File path")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'list', 'download' or 'submit' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		listHashLists()
	case "download":
		downloadCmd.Parse(os.Args[2:])
		if *downloadID == "" {
			fmt.Println("Please provide a hash list ID using -id flag.")
			os.Exit(1)
		}
		downloadHashList(*downloadID)
	case "submit":
		submitCmd.Parse(os.Args[2:])
		if *submitID == "" || *submitFile == "" {
			fmt.Println("Please provide a hash list ID using -id flag and a file path using -file flag.")
			os.Exit(1)
		}
		submitGameData(*submitID, *submitFile)
	default:
		fmt.Println("Expected 'list', 'download' or 'submit' subcommands")
		os.Exit(1)
	}
}
