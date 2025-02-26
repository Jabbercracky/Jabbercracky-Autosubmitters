// Package api provides functions for interacting with the Jabbercracky API.
package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// getJWTToken retrieves the JWT token from the environment variable
// JABBERCRACKY_API_KEY.
//
// Args:
// None
//
// Returns:
// string: The JWT token
// error: An error if the token is not set
func getJWTToken() (string, error) {
	token := os.Getenv("JABBERCRACKY_API_KEY")
	if token == "" {
		return "", fmt.Errorf("environment variable JABBERCRACKY_API_KEY is not set")
	}

	return strings.TrimSpace(token), nil
}

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
	token, err := getJWTToken()
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

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return
	}

	if hashLists, ok := result["hash_lists"].([]interface{}); ok {
		type HashList struct {
			ID   int    `json:"hash_list_id"`
			Name string `json:"hash_list_name"`
		}
		var lists []HashList

		fmt.Println("[*] Available hash lists:")
		for _, hashList := range hashLists {
			if hashListMap, ok := hashList.(map[string]interface{}); ok {
				id := int(hashListMap["hash_list_id"].(float64))
				name := hashListMap["hash_list_name"].(string)
				lists = append(lists, HashList{ID: id, Name: name})
			}
		}

		sort.Slice(lists, func(i, j int) bool {
			return lists[i].ID < lists[j].ID
		})

		for _, list := range lists {
			fmt.Printf("[*] [ID: %d] Name: %s\n", list.ID, list.Name)
		}
	} else {
		fmt.Println("Invalid response format")
	}
}

// DownloadHashList fetches the hash list with the given ID from the server
// and saves just the hash_list array to the current directory.
//
// API Endpoint: /api/game/hashlist/{id}
//
// Args:
// id (string): The ID of the hash list to download
//
// Returns:
// None
func DownloadHashList(id string) {
	token, err := getJWTToken()
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

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return
	}

	if hashList, ok := result["hash_list"].([]interface{}); ok {
		filePath := fmt.Sprintf("%s.left", id)
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		for _, hash := range hashList {
			fmt.Fprintln(file, hash)
		}
	} else {
		fmt.Println("Invalid response format")
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
	token, err := getJWTToken()
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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return
	}
	writer.Close()

	url := fmt.Sprintf("https://jabbercracky.com/api/game/submit/%s", id)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error submitting game data:", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return
	}

	if errMsg, ok := result["error"]; ok {
		fmt.Printf("[*] [ID: %s] Username: %s | Found Count: 0 | Added Score: 0 | Total Score: 0 | New Items: 0 | Error: %s\n", id, "unknown", errMsg)
		return
	}

	hashListID := result["hash_list_id"]
	username := result["username"]
	foundCount := result["found_count"]
	addedScore := result["added_score"]
	totalScore := result["total_score"]
	newItemsCount := len(result["new_items"].([]interface{}))

	fmt.Printf("[*] [ID: %s] Username: %s | Found Count: %d | Added Score: %.2f | Total Score: %.2f | New Items: %d\n",
		hashListID, username, int(foundCount.(float64)), addedScore, totalScore, newItemsCount)
}

// AutoSubmitGameData will continuously submit the file at the given path
// to the server for points every 5 minutes.
//
// The function will keep a .submitted file in the current directory to
// deduplicate submissions.
//
// API Endpoint: /api/game/hashlist/{id}
//
// Args:
// id (string): The ID of the hash list to submit
// filePath (string): The path to the hash list file
// interval (int): The interval in minutes to submit the file
//
// Returns:
// None
func AutoSubmitGameData(id string, filePath string, interval int) {
	submittedFilePath := fmt.Sprintf("%s.submitted", id)
	submittedHashes := make(map[string]bool)

	for {
		if _, err := os.Stat(submittedFilePath); os.IsNotExist(err) {
			file, err := os.Create(submittedFilePath)
			if err != nil {
				fmt.Println("Error creating submitted file:", err)
				return
			}
			file.Close()
		}

		file, err := os.Open(submittedFilePath)
		if err != nil {
			fmt.Println("Error opening submitted file:", err)
			return
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			submittedHashes[scanner.Text()] = true
		}
		file.Close()

		SubmitGameData(id, filePath)

		file, err = os.OpenFile(submittedFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening submitted file:", err)
			return
		}

		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file content:", err)
			return
		}

		hashes := strings.Split(string(fileContent), "\n")
		for _, hash := range hashes {
			if _, ok := submittedHashes[hash]; !ok {
				fmt.Fprintln(file, hash)
			}
		}
		file.Close()

		time.Sleep(time.Duration(interval) * time.Minute)
	}
}
