FROM golang:1.14.3-alpine
WORKDIR /src
ADD go.* /src/
RUN go mod download \
    && go test ./... -v -coverprofile=coverage.out -covermode=count -coverpkg=./...
