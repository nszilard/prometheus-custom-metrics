#!make
#----------------------------------------
# Settings
#----------------------------------------
.DEFAULT_GOAL := help

#--------------------------------------------------
# Variables
#--------------------------------------------------
BINARY?="prometheus-custom-metrics"
GO_FILES?=$$(find . -name '*.go')
TEST?=$$(go list ./... | grep -v /vendor/)

#--------------------------------------------------
# Targets
#--------------------------------------------------
.PHONY: bootstrap
bootstrap: ## Downloads and cleans up all dependencies
	@go mod tidy
	@go mod download
	@go install -mod=mod github.com/swaggo/swag/cmd/swag

.PHONY: fmt
fmt: ## Formats go files
	@echo "==> Formatting files..."
	@gofmt -w -s $(GO_FILES)
	@echo ""

.PHONY: check
check: ## Checks code for linting/construct errors
	@echo "==> Checking if files are well formatted..."
	@gofmt -l $(GO_FILES)
	@echo ""
	@echo "==> Checking if files pass go vet..."
	@go list -f '{{.Dir}}' ./... | xargs go vet;
	@echo ""

.PHONY: docs
docs: ## Genereates Swagger documentation
	@echo "==> Generating Swagger documentation..."
	@swag init --parseInternal

.PHONY: dev
dev: fmt check docs ## Builds a local dev version
	@go build -o ${BINARY}

.PHONY: clean
clean: ## Cleans up temporary and compiled files
	@echo "==> Cleaning up ..."
	@rm -f ${BINARY}
	@echo ""

help:
	@fgrep -h "## " $(MAKEFILE_LIST) | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-9s\033[0m %s\n", $$1, $$2}'
