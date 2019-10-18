FROM golang:latest as build
ENV GO111MODULE=off
ENV GOBUILDPATH=github.com/baxiang/hello-go
RUN mkdir -p /go/src/${GOBUILDPATH}
COPY ./ /go/src/${GOBUILDPATH}
RUN cd /go/src/${GOBUILDPATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v


FROM scratch
WORKDIR /baxiang
COPY --from=build /go/bin/hello-go ./hello-go
EXPOSE 8000
CMD ["./hello-go"]