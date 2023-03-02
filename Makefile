BIN_FILE=silly-gorest
DOCKER_REGISTRY=<specify-url>
PROJECT_NAME=silly-gorest
PROJECT_VERS=0.1


## Dependencies:
deps:
	go get github.com/gin-gonic/gin
	go get github.com/stretchr/testify

## Build:
build:	## build project and move binaries in bin/
	go build -o bin/$(BIN_FILE) main.go

clean:	## remove dependencies and build file
	go clean -i github.com/gin-gonic/gin
	go clean -i github.com/stretchr/testify
	go mod tidy
	rm -f bin/*

## Test:
test:	## run the test
	go test

## Docker:
docker-build:	## build docker image from Dockerfile
	docker build -t $(PROJECT_NAME):$(PROJECT_VERS) .

docker-release:	## publish the image to the registry
	docker tag $(PROJECT_NAME):$(PROJECT_VERS) $(DOCKER_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERS)
	docker tag $(PROJECT_NAME):$(PROJECT_VERS) $(DOCKER_REGISTRY)/$(PROJECT_NAME):latest
	docker push $(DOCKER_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERS)
	docker push $(DOCKER_REGISTRY)/$(PROJECT_NAME):latest
