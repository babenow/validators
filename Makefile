PROJECT_BIN = $(shell pwd)/bin
GOLANGCI_LINT_VERSION = v1.56.2

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
ifeq ($(OS), Windown_NT)
GOLANGCI_LINT := $(GOLANGCI_LINT).exe
endif

$(shell [ -f $(PROJECT_BIN) ] || mkdir -p $(PROJECT_BIN))

.PHONY:.install-linter
.install-linter:
	$(shell [ -f $(GOLANGCI_LINT) ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) $(GOLANGCI_LINT_VERSION))

.PHONY:lint
lint:.install-linter
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY:lint-fast
lint-fast:
	$(GOLANGCI_LINT) run ./... --fast --config=./.golangci.yml

.PHONY: test
test:lint
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := test