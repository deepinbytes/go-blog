#!/bin/bash

docker-compose up -d

docker-compose exec go-blog psql -U postgres -c "CREATE DATABASE go_blog ENCODING 'LATIN1' TEMPLATE template0 LC_COLLATE 'C' LC_CTYPE 'C';"

go get golang.org/x/tools/cmd/cover
go get github.com/mattn/goveralls
go get github.com/axw/gocov/gocov
go get -u github.com/golang/dep/cmd/dep
dep ensure

go test ./...
go run server.go