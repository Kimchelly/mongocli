# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

SOURCE_FILES?=./cmd/mongocli
BINARY_NAME=mongocli

DESTINATION=./bin/${BINARY_NAME}
INSTALL_PATH="${GOPATH}/bin/${BINARY_NAME}"

GOLANGCI_VERSION=v1.27.0
COVERAGE=coverage.out
VERSION=$(shell git describe --always --tags)
LINKER_FLAGS=-X github.com/mongodb/mongocli/internal/version.Version=${VERSION}

TEST_CMD?=go test
E2E_TAGS?=e2e

export PATH := ./bin:$(PATH)
export GO111MODULE := on

.PHONY: setup
setup:  ## Install dev tools
	@echo "==> Installing dependencies..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_VERSION)

.PHONY: link-git-hooks
link-git-hooks: ## Install git hooks
	@echo "==> Installing all git hooks..."
	find .git/hooks -type l -exec rm {} \;
	find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

.PHONY: fmt
fmt: ## Format code
	@scripts/fmt.sh

.PHONY: test
test: ## Run tests
	@echo "==> Running tests..."
	${TEST_CMD} -race -cover -count=1 -coverprofile ${COVERAGE} ./internal...

.PHONY: lint
lint: ## Run linter
	@echo "==> Linting all packages..."
	golangci-lint run

.PHONY: fix-lint
fix-lint: ## Fix linting errors
	@echo "==> Fixing lint errors"
	golangci-lint run --fix

.PHONY: check
check: test fix-lint ## Run tests and linters

.PHONY: addcopy
addcopy:
	@scripts/add-copy.sh

.PHONY: gen-mocks
gen-mocks: ## Generate mocks
	@echo "==> Generating mocks"
	go generate ./internal...

.PHONY: build
build: ## Generate a binary in ./bin
	@echo "==> Building binary"
	go build -ldflags "${LINKER_FLAGS}" -o ${DESTINATION} ${SOURCE_FILES}

.PHONY: e2e-test
e2e-test: build ## Run E2E tests
	@echo "==> Running E2E tests..."
	# the target assumes the MCLI-* environment variables are exported
	${TEST_CMD} -v -p 1 -parallel 1 -timeout 15m -tags="${E2E_TAGS}" ./e2e...

.PHONY: install
install: ## Install a binary in $GOPATH/bin
	@echo "==> Installing to $(INSTALL_PATH)"
	go install -ldflags "${LINKER_FLAGS}" ${SOURCE_FILES}
	@echo "==> Done..."

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
