# Go Blog

[![Build Status](https://travis-ci.org/deepinbytes/go-blog.svg?branch=master)](https://travis-ci.org/deepinbytes/go-blog)
[![Coverage Status](https://coveralls.io/repos/github/deepinbytes/go-blog/badge.svg?branch=master)](https://coveralls.io/github/deepinbytes/go-blog?branch=master)
[![Go Report](https://goreportcard.com/badge/github.com/deepinbytes/go-blog)](https://goreportcard.com/report/github.com/deepinbytes/go-blog)


## Packages used 

* Routing framework: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Database: [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [logrus](https://github.com/Sirupsen/logrus)
* Configuration: [viper](https://github.com/spf13/viper)
* Dependency management: [dep](https://github.com/golang/dep)
* Testing: [testify](https://github.com/stretchr/testify)


## Usage


Run `./start.sh` or manually enter the docker compose commands to bring up local server and postgresql db instance.

```
$ docker-compose build
$ docker-compose down
$ docker-compose up -d
$ docker-compose exec go-blog psql -U admin -d go_blog -a -f /docker-entrypoint-initdb.d/db.sql
```

The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

* `GET /ping`: a ping service mainly provided for health check purpose
* `GET /v1/articles`: returns a paginated list of the article
* `GET /v1/articles/:id`: returns the detailed information of an article
* `POST /v1/articles`: creates a new article
* `PUT /v1/articles/:id`: updates an existing article
* `DELETE /v1/articles/:id`: deletes an article
* `POST /v1/auth`: authenticate a user 

## Project Structure


* `models`: contains the data structures used for communication between different layers.
* `services`: contains the main business logic of the application.
* `daos`: contains the DAO (Data Access Object) layer that interacts with persistent storage.
* `apis`: contains the API layer that wires up the HTTP routes with the corresponding service APIs.

The rest of the packages are used globally:
 
* `app`: contains routing middlewares and application-level configurations
* `errors`: contains error representation and handling
* `util`: contains utility code

