LDFLAGS := -ldflags "-s -w"
GO := go
GO_INSTALELD := $(shell command -v $(GO) 2> /dev/null)
GO_NOT_INSTALLED_MESSAGE := "Go is not installed ❌\nPlease visit https://go.dev/dl/ for more information about the installation.\nAlternatively you can install go via:\n* gvm (https://github.com/moovweb/gvm)\n* asdf (https://asdf-vm.com/guide/getting-started.html)"
BUILD_OUTPUT := bring
COVERAGE_FILE := coverage.txt
CLEANUP_FILES := $(BUILD_OUTPUT) $(COVERAGE_FILE)

.PHONY: setup
setup:
ifdef GO_INSTALELD
	@echo "You are good to go 🚀"
else
	@echo $(GO_NOT_INSTALLED_MESSAGE)
endif

.PHONY: build
build:
	$(GO) build -o $(BUILD_OUTPUT) $(LDFLAGS) .

.PHONY: test
test: $(TEST_MODULES)
	go test -v -short -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...

.PHONY: install
install:
	go install $(LDFLAGS) .

.PHONY: clean
clean:
	rm -f $(CLEANUP_FILES)
