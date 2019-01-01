# Project Name

SHA1         		:= $(shell git rev-parse --verify --short HEAD)
BUILD_IMAGE			:= $(shell echo "golang:1.11.4")
PWD					:= $(shell pwd)

.DEFAULT_GOAL := test

.PHONY: test
test:
	@echo "Running Tests"

	docker run -e GO111MODULE=auto --rm -t -v $(PWD):/usr/src/myapp -w /usr/src/myapp $(BUILD_IMAGE) sh -c "go test -cover -v ./... -count=1"
	@echo "Completed tests"


