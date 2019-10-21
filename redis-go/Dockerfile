FROM golang:latest as build
WORKDIR /go/src/github.com/baxiang/redis-go
COPY . .
RUN cd /go/src/github.com/baxiang/redis-go && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v

FROM ubuntu:18.04
WORKDIR /app
COPY --from=build /go/bin/redis-go .
EXPOSE 8090
CMD ["./redis-go"]