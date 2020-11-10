.PHONY: build
build:
	go build -buildmode plugin -o go-kong.so main.go

.PHONY: docker
docker:
	docker build -t kong-demo .
	docker run -it --rm --name kong-plugins -p 8888:8000 kong-demo
