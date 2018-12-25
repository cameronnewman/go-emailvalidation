# Project Name

SHA1         		:= $(shell git rev-parse --verify --short HEAD)
BUILD_IMAGE			:= $(shell echo "golang:1.11.4")
PWD					:= $(shell pwd)


.PHONY: tests
tests:
	@echo "Running Tests"

	docker run --rm -t -v $(PWD):/go/src -w /go/src $(BUILD_IMAGE) go test -cover -v $(go list ./... | grep -v /example/)
	@echo "Completed tests"


