# Jabbercracky Autosubmitters

## Getting Started Guide
The following guide will help you set up and run the Jabbercracky
Autosubmitters. The autosubmitters are a set of scripts that are written in
various languages that will automatically submit your Jabbercracky founds to
the server every hour.

### Prerequisites
- Python 3.6 or higher

### Installation
We will be using the Python autosubmitter for this guide. To install the Python
script, download the `autosubmitter.py` file from the repository and place it
into a directory of your choice.

```bash
wget https://raw.githubusercontent.com/Jabbercracky/Jabbercracky-Autosubmitters/refs/heads/main/python/jabbercracky-client.py
```

### Configuration
Before running the autosubmitter, you will need to get your API key from the
Jabbercracky website. To do this, visit the [Jabbercracky
website](https://jabbercracky.com) and sign in. Once you are signed in, click
on the "Game" tab and then click on the "Account" tab. Under "Copy
Authentication Token", click the "Copy" button to copy your API key to your
clipboard. Please keep this key safe and do not share it with anyone as it can
be used to log in to your account.

Once you have your API key, you can set it to the `JABBERCRACKY_API_KEY`
environment variable. You can do this by running the following command:

```bash
export JABBERCRACKY_API_KEY="your-api"
```

or by adding the following line to your `.bashrc` or `.zshrc` file:

```bash
export JABBERCRACKY_API_KEY="your-api"
```

or by calling the variable before running the script:

```bash
JABBERCRACKY_API_KEY="your-api" python3 autosubmitter.py
```

### Running the Autosubmitter
To run the autosubmitter, simply run the following command:

```bash
python3 autosubmitter.py list
```

This will list all available lists to work on. To download a list, run the
`download` with the `-id` flag and the list ID:

```bash
python3 autosubmitter.py download -id 1
```

To submit your founds to the server, run the `submit` command with the `-id`
flag, the list ID, and the `-file` flag with the path to the founds file:

```bash
python3 autosubmitter.py submit -id 1 -file founds.txt
```

Then to automatically submit every hour, run the `autosubmit` command.

```bash
python3 autosubmitter.py autosubmit -id 1 -file founds.txt
```

This will also create a file called `<id>.submitted` in the same directory as
the script to keep track of the founds that have been submitted.
