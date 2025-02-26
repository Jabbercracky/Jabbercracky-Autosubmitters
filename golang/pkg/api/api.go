// Package api provides functions for interacting with the Jabbercracky API.
package api

import (
	"fmt"
	"io/ioutil"
	"jabbercracky-api-client/pkg/utils"
	"net/http"
	"os"
)

// ListHashLists fetches the list of hash lists from the server
// and prints it to the console.
//
// API Endpoint: /api/game/hashlist
//
// Args:
// None
//
// Returns:
// None
func ListHashLists() {
	token, err := utils.GetJWTToken()
	if err != nil {
		fmt.Println("Error getting JWT token:", err)
		return
	}

	req, err := http.NewRequest("GET", "https://jabbercracky.com/api/game/hashlist", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

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

// DownloadHashList fetches the hash list with the given ID from the server
// and saves it to the current directory.
//
// API Endpoint: /api/game/hashlist/{id}
//
// Args:
// id (string): The ID of the hash list to download
//
// Returns:
// None
func DownloadHashList(id string) {
	token, err := utils.GetJWTToken()
	if err != nil {
		fmt.Println("Error getting JWT token:", err)
		return
	}

	url := fmt.Sprintf("https://jabbercracky.com/api/game/hashlist/%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

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

	filePath := fmt.Sprintf("%s.left", id)
	err = ioutil.WriteFile(filePath, body, 0644)
	if err != nil {
		fmt.Println("Error saving hash list:", err)
		return
	}
}

// SubmitGameData submits the hash list with the given ID to the server
// for points.
//
// API Endpoint: /api/game/hashlist/{id}
//
// Args:
// id (string): The ID of the hash list to submit
// filePath (string): The path to the hash list file
//
// Returns:
// None
func SubmitGameData(id string, filePath string) {
	token, err := utils.GetJWTToken()
	if err != nil {
		fmt.Println("Error getting JWT token:", err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	url := fmt.Sprintf("https://jabbercracky.com/api/game/submit/%s", id)
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)

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
