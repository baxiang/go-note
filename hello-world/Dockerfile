FROM registry.cn-beijing.aliyuncs.com/baxiang/golang:1.14 as build
ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
ENV GOBUILDPATH=github.com/baxiang/hello-world
RUN mkdir -p /go/src/${GOBUILDPATH}
COPY ./ /go/src/${GOBUILDPATH}
RUN cd /go/src/${GOBUILDPATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v


FROM alpine as dev
WORKDIR /hello-wordld
COPY --from=build /go/bin/hello-world ./hello-world
EXPOSE 8080
CMD ["./hello-world"]