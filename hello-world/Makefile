DOCKER_REGISTRY  = registry.cn-beijing.aliyuncs.com/baxiang
DOCKER_IMAGE     = hello-world
DOCKER_FULL_NAME = $(DOCKER_REGISTRY)/$(DOCKER_IMAGE)
DOCKER_TAG = v1.0



.PHONY: docker
docker:
	@docker build --target dev  -t $(DOCKER_FULL_NAME):$(DOCKER_TAG) -f Dockerfile .
	@docker push $(DOCKER_FULL_NAME):$(DOCKER_TAG)




