GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
BINARY_NAME=main
VERSION?=0.0.0#Docker image release version
EXPORT_RESULT?=false# for CI please set EXPORT_RESULT to true
TEST?=false#is the container test?

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build coverage lint

all: help

## Generate:
proto-generate: ## Protobuf Generate
	protoc -I=../op-proto/pkg/proto/model \
		--go_out=../../../ \
		--go-grpc_out=../../../ \
		../op-proto/pkg/proto/model/button.proto \
		../op-proto/pkg/proto/model/led.proto \
		../op-proto/pkg/proto/model/tick.proto \
		../op-proto/pkg/proto/model/dlr.proto \
		--experimental_allow_proto3_optional

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)