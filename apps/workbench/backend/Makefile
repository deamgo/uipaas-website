.PHONY: help
help:  ## Display this help message.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: test
test: mod fmt lint ## Run unit tests after formatting and linting
	@GOARCH=amd64 go test -count=1 ./... -ldflags="-s=false" -gcflags="-l" --cover

.PHONY: fmt
fmt: ## Run go fmt for code formatting
	@echo "Running go fmt"
	@go fmt ./...

.PHONY: lint
lint: ## Run golangci-lint
	@echo "Running golangci-lint"
	@golangci-lint run ./...

.PHONY: mod
mod: ## go mod tidy
	@echo "Running go mod tidy"
	@go mod tidy
