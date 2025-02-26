# Jabbercracky Golang Autosubmitter

## Overview

This is a command line interface for submitting game data to the Jabbercracky
API. The CLI provides several subcommands for listing available hash lists,
downloading hash lists, submitting game data, and auto-submitting game data.

## Prerequisites

- Go 1.16 or higher
- A valid Jabbercracky API key

## Setting Up

Before running any scripts, you need to set your Jabbercracky API key in the environment variable `JABBERCRACKY_API_KEY`.

```sh
export JABBERCRACKY_API_KEY="your_api_key"
```

This can be done in your shell profile or in a script that you run before
executing the CLI. Alternatively, you can set the environment variable inline when running the script:
```sh
JABBERCRACKY_API_KEY="your_api_key" jabbercracky-client list
```

## Usage

The command line interface provides several subcommands. Here are the available options:

- list: Lists all available hash lists.
- download: Downloads the hash list specified by the given ID.
- submit: Submits game data from the specified file to the hash list with the given ID.
- auto-submit: Automatically submits game data from the specified file to the hash list with the given ID every 5 minutes.

### Examples
List Hash Lists
```sh
jabbercracky-client list
```

Download a Hash List
```sh
jabbercracky-client download -id 12345
```

Submit Game Data
```sh
jabbercracky-client submit -id 12345 -file path/to/file.txt
```

Auto-Submit Game Data
```sh
jabbercracky-client auto-submit -id 12345 -file path/to/file.txt
```

