#!make
#----------------------------------------
# Settings
#----------------------------------------
.DEFAULT_GOAL := help

#--------------------------------------------------
# Variables
#--------------------------------------------------
BINARY?="prometheus-custom-metrics"
TEST?=$$(go list ./...)
GO_FILES?=$$(find . -name '*.go')

#--------------------------------------------------
# Targets
#--------------------------------------------------
bootstrap: ## Downloads and cleans up all dependencies
	@go mod tidy
	@go mod download

fmt: ## Formats go files
	@echo "==> Formatting files..."
	@gofmt -w $(GO_FILES)
	@echo ""

check: ## Checks code for linting/construct errors
	@echo "==> Checking if files are formatted..."
	@gofmt -l $(GO_FILES)
	@echo "    [✓]\n"
	@echo "==> Checking if files pass go vet..."
	@go list -f '{{.Dir}}' ./... | xargs go vet;
	@echo "    [✓]\n"

test: check ## Runs all tests
	@echo "==> Running tests..."
	@go test $(TEST) -parallel=20
	@echo ""

coverage: ## Runs code coverage
	@go test $(TEST) -race -coverprofile=.target/coverage.out -covermode=atomic

show-coverage: coverage ## Shows code coverage report in your web browser
	@go tool cover -html=.target/coverage.out

dev: check ## Builds a local dev version
	@go build

package: clean bootstrap check test ## Packages the binary for release
	@mkdir -p .target/bin
	@env GOOS=linux GOARCH=amd64 go build -o .target/bin/${BINARY}
	@docker build -t nszilard/prometheus-custom-metrics .

.PHONY: bootstrap check package fmt test coverage show-coverage clean help

clean: ## Cleans up temporary and compiled files
	@echo "==> Cleaning up ..."
	@rm -rf .target/*
	@echo "    [✓]"
	@echo ""

help: ## Shows available targets
	@fgrep -h "## " $(MAKEFILE_LIST) | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-13s\033[0m %s\n", $$1, $$2}'
