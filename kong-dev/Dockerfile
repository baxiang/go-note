FROM kong/go-plugin-tool:latest-centos-7 as builder
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir -p /go/src/kong-plugins&&mkdir -p /go/src/go-pluginserver
RUN git clone https://github.com/Kong/go-pluginserver /go/src/go-pluginserver \
&& cd /go/src/go-pluginserver && make

COPY ./ /go/src/kong-plugins
RUN cd /go/src/kong-plugins&& make build

FROM kong:2.2.0-centos

ENV KONG_DATABASE=off
ENV KONG_GO_PLUGINS_DIR=/tmp/go-plugins
ENV KONG_DECLARATIVE_CONFIG=/tmp/config.yml
ENV KONG_PLUGINS=go-kong
COPY --from=builder /go/src/go-pluginserver/go-pluginserver /usr/local/bin/go-pluginserver
RUN mkdir /tmp/go-plugins
COPY --from=builder /go/src/kong-plugins/go-kong.so /tmp/go-plugins/go-kong.so
COPY config.yml /tmp/config.yml
