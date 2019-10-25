FROM golang:latest as build
WORKDIR /go/src/github.com/baxiang/mysql-go
COPY . .
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v

FROM scratch
WORKDIR /app
COPY --from=build /go/bin/mysql-go .
CMD ["./mysql-go"]