GOBASE=$(shell pwd)
CGO_ENABLED=0
PROJECT_NAME=$(shell basename "$(PWD)")
GOPROXY=https://goproxy.cn,direct

fmt:
	@go fmt

test:
	ROOTDIR=$(GOBASE) ENV=test go test -v ./... -cover

clean:
	@echo "clean all cache"
	@rm -rf $(PROJECT_NAME)
	@rm -rf ./vendor

install:
	@echo "set module on"
	export GO111MODULE=on
	@echo "set golang proxy"
	export GOPROXY=$(GOPROXY)
	@echo "install deps"
	@go mod tidy
	@echo "verify deps"
	@go mod verify
	@go mod vendor

build: clean
	@echo "build project"
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -o $(PROJECT_NAME)