# Jabbercracky Autosubmitters

This repository contains `curl` examples for several API endpoints.

# /api/game/submit/<id>
```sh
$ curl "https://jabbercracky.com/api/game/submit/[TARGET_HASHLIST]" -X POST
-H "Content-Type: multipart/form-data"
-H "Cookie: auth_token=[YOUR JWT HERE]"
--form "file=@found_hashes_list_1.txt"
```
