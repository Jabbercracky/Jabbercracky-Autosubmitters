// Package utils contains utility functions that are used by the main package.
package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// GetJWTToken reads the JWT token from the user's home directory.
//
// The token is stored in a file named .jabbercracky in the user's home
// directory.
//
// Args:
// None
//
// Returns:
// string: The JWT token
// error: An error if one occurred
func GetJWTToken() (string, error) {
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
