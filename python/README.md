# Jabbercracky Python Autosubmitter

## Overview

This is a command line interface for submitting game data to the Jabbercracky
API. The CLI provides several subcommands for listing available hash lists,
downloading hash lists, submitting game data, and auto-submitting game data.

## Prerequisites

- Python 3.6 or higher
- A valid Jabbercracky API key

## Setting Up

Before running any scripts, you need to set your Jabbercracky API key in the environment variable `JABBERCRACKY_API_KEY`.

```sh
export JABBERCRACKY_API_KEY="your_api_key"
```

This can be done in your shell profile or in a script that you run before
executing the CLI. Alternatively, you can set the environment variable inline when running the script:
```sh
JABBERCRACKY_API_KEY="your_api_key" python3 jabbercracky-client.py list
```

## Usage

The command line interface provides several subcommands. Here are the available options:

- list: Lists all available hash lists.
- download: Downloads the hash list specified by the given ID.
- submit: Submits game data from the specified file to the hash list with the given ID.
- auto-submit: Automatically submits game data from the specified file to the hash list with the given ID every hour.

### Examples
List Hash Lists
```sh
python3 jabbercracky-client.py list
```

Download a Hash List
```sh
python3 jabbercracky-client.py download -id 12345
```

Submit Game Data
```sh
python3 jabbercracky-client.py submit -id 12345 -file path/to/file.txt
```

Auto-Submit Game Data
```sh
python3 jabbercracky-client.py auto-submit -id 12345 -file path/to/file.txt
```
