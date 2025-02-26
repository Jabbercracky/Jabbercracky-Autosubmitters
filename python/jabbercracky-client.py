import argparse
import os
import requests
import time
import json

def get_jwt_token():
    """
    Retrieves the JWT token from the environment variable JABBERCRACKY_API_KEY.

    Returns:
    str: The JWT token
    """
    token = os.getenv('JABBERCRACKY_API_KEY')
    if not token:
        raise EnvironmentError("Environment variable JABBERCRACKY_API_KEY is not set")
    return token.strip()

def list_hash_lists():
    """
    Fetches the list of hash lists from the server and prints it to the console.

    API Endpoint: /api/game/hashlist
    """
    token = get_jwt_token()
    headers = {'Authorization': f'Bearer {token}'}
    response = requests.get('https://jabbercracky.com/api/game/hashlist', headers=headers)
    if response.status_code != 200:
        print("Error fetching hash lists:", response.text)
        return

    result = response.json()
    hash_lists = result.get('hash_lists', [])

    print("[*] Available hash lists:")
    for hash_list in sorted(hash_lists, key=lambda x: x['hash_list_id']):
        print(f"[*] [ID: {hash_list['hash_list_id']}] Name: {hash_list['hash_list_name']}")

def download_hash_list(id):
    """
    Fetches the hash list with the given ID from the server and saves just the hash_list array to the current directory.

    API Endpoint: /api/game/hashlist/{id}

    Args:
    id (str): The ID of the hash list to download
    """
    token = get_jwt_token()
    headers = {'Authorization': f'Bearer {token}'}
    response = requests.get(f'https://jabbercracky.com/api/game/hashlist/{id}', headers=headers)
    if response.status_code != 200:
        print("Error fetching hash list:", response.text)
        return

    result = response.json()
    hash_list = result.get('hash_list', [])

    with open(f"{id}.left", 'w') as file:
        for hash_item in hash_list:
            file.write(f"{hash_item}\n")

def submit_game_data(id, file_path):
    """
    Submits the hash list with the given ID to the server for points.

    API Endpoint: /api/game/hashlist/{id}

    Args:
    id (str): The ID of the hash list to submit
    file_path (str): The path to the hash list file
    """
    token = get_jwt_token()
    headers = {'Authorization': f'Bearer {token}'}

    with open(file_path, 'rb') as file:
        files = {'file': file}
        response = requests.post(f'https://jabbercracky.com/api/game/submit/{id}', headers=headers, files=files)
    
    if response.status_code != 200:
        print("Error submitting game data:", response.text)
        return

    result = response.json()
    if 'error' in result:
        print(f"[*] [ID: {id}] Username: unknown | Found Count: 0 | Added Score: 0 | Total Score: 0 | New Items: 0 | Error: {result['error']}")
        return

    print(f"[*] [ID: {result['hash_list_id']}] Username: {result['username']} | Found Count: {result['found_count']} | Added Score: {result['added_score']} | Total Score: {result['total_score']} | New Items: {len(result['new_items'])}")

def auto_submit_game_data(id, file_path, interval):
    """
    Continuously submits the file at the given path to the server for points every 5 minutes.

    The function will keep a .submitted file in the current directory to deduplicate submissions.

    API Endpoint: /api/game/hashlist/{id}

    Args:
    id (str): The ID of the hash list to submit
    file_path (str): The path to the hash list file
    interval (int): The interval in minutes to submit the file
    """
    submitted_file_path = f"{id}.submitted"
    submitted_hashes = set()

    while True:
        if not os.path.exists(submitted_file_path):
            open(submitted_file_path, 'w').close()

        with open(submitted_file_path, 'r') as file:
            for line in file:
                submitted_hashes.add(line.strip())

        submit_game_data(id, file_path)

        with open(submitted_file_path, 'a') as file:
            with open(file_path, 'r') as data_file:
                for line in data_file:
                    hash_item = line.strip()
                    if hash_item not in submitted_hashes:
                        file.write(f"{hash_item}\n")
                        submitted_hashes.add(hash_item)

        time.sleep(interval * 60)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='jabbercracky-client command line interface')
    subparsers = parser.add_subparsers(dest='command')

    list_parser = subparsers.add_parser('list', help='Lists all available hash lists')
    download_parser = subparsers.add_parser('download', help='Downloads the hash list specified by the given ID')
    download_parser.add_argument('-id', required=True, help='Hash List ID')

    submit_parser = subparsers.add_parser('submit', help='Submits game data from the specified file to the hash list with the given ID')
    submit_parser.add_argument('-id', required=True, help='Hash List ID')
    submit_parser.add_argument('-file', required=True, help='File path')

    auto_submit_parser = subparsers.add_parser('auto-submit', help='Automatically submits game data from the specified file to the hash list with the given ID every 5 minutes')
    auto_submit_parser.add_argument('-id', required=True, help='Hash List ID')
    auto_submit_parser.add_argument('-file', required=True, help='File path')

    args = parser.parse_args()

    if args.command == 'list':
        list_hash_lists()
    elif args.command == 'download':
        download_hash_list(args.id)
    elif args.command == 'submit':
        submit_game_data(args.id, args.file)
    elif args.command == 'auto-submit':
        auto_submit_game_data(args.id, args.file, 5)
    else:
        parser.print_help()
