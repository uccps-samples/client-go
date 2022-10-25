all: build
.PHONY: all

build:
	go build github.com/uccps-samples/client-go/...
.PHONY: build

build-examples:
	go build -o examples/build/app github.com/uccps-samples/client-go/examples/build/...
.PHONY: build-examples

verify:
	hack/verify-codegen.sh
.PHONY: verify

generate:
	hack/update-codegen.sh
.PHONY: generate
