#!/bin/bash

docker compose exec app go test ./...

curl -X GET 'http://localhost:1323/get_friend_list?id=1'
curl -X GET 'http://localhost:1323/get_friend_of_friend_list?id=1'
curl -X GET 'http://localhost:1323/get_friend_of_friend_list_paging?id=1&page=1&limit=1'
