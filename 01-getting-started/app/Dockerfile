FROM golang:1-alpine

WORKDIR /go/src/github.com/shirokuma-software/advanced-kubernetes-examples
CMD ["./app"]
COPY main.go .
RUN go build -o app main.go
