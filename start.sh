#!/bin/sh
docker-compose build

docker-compose down

docker-compose up

docker-compose exec go-blog psql -U admin -d go_blog -a -f /docker-entrypoint-initdb.d/db.sql
