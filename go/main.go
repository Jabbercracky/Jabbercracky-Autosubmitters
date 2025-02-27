// Package main provides a command line interface for the
// jabbercracky-api-client package.
package main

import (
	"flag"
	"fmt"
	"os"

	"jabbercracky-client/pkg/api"
)

func main() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.Usage = func() {
		fmt.Println("[*] Usage: list")
		fmt.Println("  Lists all available hash lists.")
		fmt.Println("Example:")
		fmt.Println("  $ jabbercracky-client list")
	}

	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
	downloadCmd.Usage = func() {
		fmt.Println("[*] Usage: download -id <Hash List ID>")
		fmt.Println("  Downloads the hash list specified by the given ID.")
		fmt.Println("Example:")
		fmt.Println("  $ jabbercracky-client download -id 12345")
	}
	downloadID := downloadCmd.String("id", "", "Hash List ID")

	submitCmd := flag.NewFlagSet("submit", flag.ExitOnError)
	submitCmd.Usage = func() {
		fmt.Println("[*] Usage: submit -id <Hash List ID> -file <File Path>")
		fmt.Println("  Submits game data from the specified file to the hash list with the given ID.")
		fmt.Println("Example:")
		fmt.Println("  $ jabbercracky-client submit -id 12345 -file path/to/file.txt")
	}
	submitID := submitCmd.String("id", "", "Hash List ID")
	submitFile := submitCmd.String("file", "", "File path")

	autoSubmitCmd := flag.NewFlagSet("auto-submit", flag.ExitOnError)
	autoSubmitCmd.Usage = func() {
		fmt.Println("[*] Usage: auto-submit -id <Hash List ID> -file <File Path>")
		fmt.Println("  Automatically submits game data from the specified file to the hash list with the given ID every 5 minutes.")
		fmt.Println("Example:")
		fmt.Println("  $ jabbercracky-client auto-submit -id 12345 -file path/to/file.txt")
	}
	autoSubmitID := autoSubmitCmd.String("id", "", "Hash List ID")
	autoSubmitFile := autoSubmitCmd.String("file", "", "File path")

	if len(os.Args) < 2 {
		fmt.Println("[*] Available flags:")
		fmt.Println("[*] The client will look for credentials in the environment variable JABBERCRACKY_API_KEY.")
		fmt.Println()
		listCmd.Usage()
		downloadCmd.Usage()
		fmt.Println()
		submitCmd.Usage()
		fmt.Println()
		autoSubmitCmd.Usage()
		fmt.Println()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		api.ListHashLists()
	case "download":
		downloadCmd.Parse(os.Args[2:])
		if *downloadID == "" {
			fmt.Println("[!] Please provide a hash list ID using -id flag.")
			downloadCmd.Usage()
			os.Exit(1)
		}
		api.DownloadHashList(*downloadID)
	case "submit":
		submitCmd.Parse(os.Args[2:])
		if *submitID == "" || *submitFile == "" {
			fmt.Println("[!] Please provide a hash list ID using -id flag and a file path using -file flag.")
			submitCmd.Usage()
			os.Exit(1)
		}
		api.SubmitGameData(*submitID, *submitFile)
	case "auto-submit":
		autoSubmitCmd.Parse(os.Args[2:])
		if *autoSubmitID == "" || *autoSubmitFile == "" {
			fmt.Println("[!] Please provide a hash list ID using -id flag and a file path using -file flag.")
			autoSubmitCmd.Usage()
			os.Exit(1)
		}
		api.AutoSubmitGameData(*autoSubmitID, *autoSubmitFile, 5)
	default:
		fmt.Println("[*] Available flags:")
		fmt.Println("[*] The client will look for credentials in the environment variable JABBERCRACKY_API_KEY.")
		fmt.Println()
		listCmd.Usage()
		fmt.Println()
		downloadCmd.Usage()
		fmt.Println()
		submitCmd.Usage()
		fmt.Println()
		autoSubmitCmd.Usage()
		fmt.Println()
		os.Exit(1)
	}
}
