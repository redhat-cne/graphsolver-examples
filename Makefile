.PHONY: all clean test build
lint:
	golangci-lint run
# Install golangci-lint	
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GO_PATH}/bin ${GOLANGCI_VERSION}
build:
	go build simplegraphsolver.go
build-all:
	$(MAKE) build -C .
	$(MAKE) build -C ./examples/basic

lint-all:
	$(MAKE) lint -C .
	$(MAKE) lint -C ./pkg/lib
	$(MAKE) lint -C ./pkg/export
	$(MAKE) lint -C ./examples/basic
vet:
	go vet ${GO_PACKAGES}
test:
	./scripts/test.sh
