#! /bin/bash

docker-compose up -d
go test ./...
docker-compose down
