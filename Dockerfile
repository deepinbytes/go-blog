FROM golang:latest

RUN mkdir -p /go/src/github.com/deepinbyes/go-blog
WORKDIR /go/src/github.com/deepinbyes/go-blog
# Copy the local package files to the container’s workspace.
ADD . .

# Install our dependencies

RUN go get -u github.com/golang/dep/cmd/dep
COPY ./Gopkg.toml /go/src/github.com/deepinbyes/go-blog
RUN dep ensure -v

# Build tbe binary and give permissions
RUN go build server.go
RUN chmod 777 server

# Set binary as entrypoint
ENTRYPOINT ./server

# Expose default port (8080)
EXPOSE 8080

CMD ./server