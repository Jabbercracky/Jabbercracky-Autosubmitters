// Package main provides a command line interface for the
// jabbercracky-api-client package.
package main

import (
	"flag"
	"fmt"
	"jabbercracky-api-client/pkg/api"
	"os"
)

func main() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
	submitCmd := flag.NewFlagSet("submit", flag.ExitOnError)
	autoSubmitCmd := flag.NewFlagSet("auto-submit", flag.ExitOnError)

	downloadID := downloadCmd.String("id", "", "Hash List ID")
	submitID := submitCmd.String("id", "", "Hash List ID")
	submitFile := submitCmd.String("file", "", "File path")
	autoSubmitID := autoSubmitCmd.String("id", "", "Hash List ID")
	autoSubmitFile := autoSubmitCmd.String("file", "", "File path")

	if len(os.Args) < 2 {
		fmt.Println("[!] Expected one of the following subcommands: 'list', 'download', 'submit, 'auto-submit'")
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
			os.Exit(1)
		}
		api.DownloadHashList(*downloadID)
	case "submit":
		submitCmd.Parse(os.Args[2:])
		if *submitID == "" || *submitFile == "" {
			fmt.Println("[!] Please provide a hash list ID using -id flag and a file path using -file flag.")
			os.Exit(1)
		}
		api.SubmitGameData(*submitID, *submitFile)
	case "auto-submit":
		autoSubmitCmd.Parse(os.Args[2:])
		if *autoSubmitID == "" || *autoSubmitFile == "" {
			fmt.Println("[!] Please provide a hash list ID using -id flag and a file path using -file flag.")
			os.Exit(1)
		}
		api.AutoSubmitGameData(*autoSubmitID, *autoSubmitFile, 5)
	default:
		fmt.Println("[!] Expected one of the following subcommands: 'list', 'download', 'submit, 'auto-submit'")
		os.Exit(1)
	}
}
