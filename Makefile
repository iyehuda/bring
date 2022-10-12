LDFLAGS := -ldflags "-s -w"
GO := go
GO_INSTALELD := $(shell command -v $(GO) 2> /dev/null)
GO_NOT_INSTALLED_MESSAGE := "Go is not installed ❌\nPlease visit https://go.dev/dl/ for more information about the installation.\nAlternatively you can install go via:\n* gvm (https://github.com/moovweb/gvm)\n* asdf (https://asdf-vm.com/guide/getting-started.html)"
LINT := golangci-lint
LINT_INSTALELD := $(shell command -v $(LINT) 2> /dev/null)
LINT_NOT_INSTALLED_MESSAGE := "$(LINT) is not installed ❌\nPlease visit https://golangci-lint.run/usage/install/ for more information about the installation.\nAlternatively you can install $(LINT) via running:\n$$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.50.0"
ENTRYPOINT := ./cmd/bring
BUILD_OUTPUT := bring
COVERAGE_FILE := coverage.txt
CLEANUP_FILES := $(BUILD_OUTPUT) $(COVERAGE_FILE)
TEST_PARALLEL := 10

.PHONY: all
all: lint test integration-tests build install

.PHONY: setup
setup:
ifndef GO_INSTALELD
	@echo $(GO_NOT_INSTALLED_MESSAGE)
else ifndef LINT_INSTALELD
	@echo $(LINT_NOT_INSTALLED_MESSAGE)
else
	@echo "You are good to go 🚀"
endif

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	$(GO) build -o $(BUILD_OUTPUT) $(LDFLAGS) $(ENTRYPOINT)

.PHONY: test
test:
	$(GO) test -v -parallel $(TEST_PARALLEL) ./...

.PHONY: integration-tests
integration-tests:
	$(GO) test -v -parallel $(TEST_PARALLEL) -tags=integration ./integration/...

.PHONY: install
install:
	$(GO) install $(LDFLAGS) $(ENTRYPOINT)

.PHONY: clean
clean:
	rm -f $(CLEANUP_FILES)
