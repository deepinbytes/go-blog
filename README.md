# Go Blog

[![Build Status](https://travis-ci.org/deepinbytes/go-blog.svg?branch=master)](https://travis-ci.org/deepinbytes/go-blog)
[![Coverage Status](https://coveralls.io/repos/github/deepinbytes/go-blog/badge.svg?branch=master)](https://coveralls.io/github/deepinbytes/go-blog?branch=master)
[![Go Report](https://goreportcard.com/badge/github.com/deepinbytes/go-blog)](https://goreportcard.com/report/github.com/deepinbytes/go-blog)


#### Packages used 

* Routing framework: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Database: [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [logrus](https://github.com/Sirupsen/logrus)
* Configuration: [viper](https://github.com/spf13/viper)
* Dependency management: [dep](https://github.com/golang/dep)
* Testing: [testify](https://github.com/stretchr/testify)


## Usage

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires Go 1.5 or above.

After installing Go, run the following commands to download and install the dependencies:

```shell
# install dep
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# fetch the dependent packages
cd $GOPATH/deepinbytes/go-blog
dep ensure
```

Run docker to bring up local `postgresql` db instance.

```
$ docker-compose build
$ docker-compose up -d
```

Next, create the database named `go_blog` by executing the command:

```
docker exec -it go-blog psql -U postgres -c "CREATE DATABASE go_blog ENCODING 'LATIN1' TEMPLATE template0 LC_COLLATE 'C' LC_CTYPE 'C';"
```

Run the unit tests by using the following command from the project directory:
```
go test ./...
```


Now you can build and run the application by running the following command under the
`$GOPATH/deepinbytes/go-blog` directory:

```shell
go run server.go
```

or simply the following if you have the `make` tool:

```shell
make
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

