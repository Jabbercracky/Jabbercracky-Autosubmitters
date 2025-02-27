# Jabbercracky Autosubmitters

This repository contains `curl` examples for several API endpoints. Note that these are single HTTP requests unlike other scripts.

## /api/game/submit/\<id>
```sh
$ curl "https://jabbercracky.com/api/game/submit/[TARGET_HASHLIST]" -X POST
-H "Content-Type: multipart/form-data"
-H "Cookie: auth_token=[YOUR JWT HERE]"
--form "file=@found_hashes_list_1.txt"
```
## /api/game/hashlist
```sh
$ curl "https://jabbercracky.com/api/game/hashlist" -X GET \
-H "Authorization: Bearer [YOUR JWT HERE]"
```

## /api/game/hashlist/\<id>
```sh
$ curl "https://jabbercracky.com/api/game/hashlist/[HASHLIST_ID]" -X GET \
-H "Authorization: Bearer [YOUR JWT HERE]"
```
