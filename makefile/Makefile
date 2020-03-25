APPNAME="main"
.PHONY: build
build:clean
	@go build -o ${APPNAME} main.go
.PHONY: clean
clean:
	@rm -rf ${APPNAME}
.PHONY: run
run:
	@rm -rf ${APPNAME}
	@go build -o ${APPNAME} main.go
	@./${APPNAME}
.PHONE: build-hello
build-hello:
	@${MAKE} -C hello build
