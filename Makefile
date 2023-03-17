help: HELP_SCRIPT = \
	if (/^([a-zA-Z0-9-\.\/]+).*?: description\s*=\s*(.+)/) { \
		printf "\033[34m%-40s\033[0m %s\n", $$1, $$2 \
	} elsif(/^\#\#\#\s*(.+)/) { \
		printf "\033[33m>> %s\033[0m\n", $$1 \
	}

.PHONY: help
help:
	@perl -ne '$(HELP_SCRIPT)' $(MAKEFILE_LIST)

.PHONY: setup/darwin
setup/darwin: description = Install protobuf tooling for macOS
setup/darwin:
	# Protocol compiler and linter
	brew tap yoheimuta/protolint
	brew install protobuf@3.11.4 protolint

	# Go plugin used by the protocol compiler
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: setup/linux
setup/linux: description = Install protobuf tooling for linux
setup/linux:
	# Protocol compiler
	PROTOC_ZIP=protoc-3.10.1-linux-x86_64.zip
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.10.1/$PROTOC_ZIP
	sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
	sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
	rm -f $PROTOC_ZIP

	# Go plugin used by the protocol compiler
	go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: generate/go
generate/go: description = Compile protobuf schemas for Go
generate/go: clean-go
generate/go:
	mkdir -p build/go/protos
	mkdir -p build/go/protos/common
	mkdir -p build/go/protos/records
	mkdir -p build/go/protos/events
	mkdir -p build/go/protos/services

	docker run --platform linux/amd64 --rm -w $(PWD) -v $(PWD):/defs -w${PWD} namely/protoc-all:1.51_1 \
		--go-source-relative \
		-l go \
		-i /defs/protos \
		-i /defs/protos/common \
		-i /defs/protos/records \
		-i /defs/protos/events \
		-d /defs/protos/services \
		-o /defs/build/go/protos \
		protos/services/*.proto


	docker run --platform linux/amd64 --rm -w $(PWD) -v $(PWD):/defs -w${PWD} namely/protoc-all:1.51_1 \
		-d /defs/protos/common \
		--go-source-relative \
		-l go \
		-o /defs/build/go/protos/common \
		protos/common/*.proto

	docker run --platform linux/amd64 --rm -w $(PWD) -v $(PWD):/defs -w${PWD} namely/protoc-all:1.51_1 \
		-d /defs/protos/records \
		--go-source-relative \
		-l go \
		-o /defs/build/go/protos/records \
		protos/records/*.proto

	docker run --platform linux/amd64 --rm -w $(PWD) -v $(PWD):/defs -w${PWD} namely/protoc-all:1.51_1 \
		-d /defs/protos/events \
		--go-source-relative \
		-l go \
		-o /defs/build/go/protos/events \
		protos/events/*.proto

# Protoset files contain binary encoded google.protobuf.FileDescriptorSet protos
.PHONY: generate/protoset
generate/protoset: description = Generate protoset for services
generate/protoset:
	mkdir -p build/go/protoset/
	docker run --platform linux/amd64 -w $(PWD) -v $(PWD):/defs namely/protoc-all:1.51_1 \
		-i /defs/protos \
		-i /defs/protos/common \
		-i /defs/protos/records \
		-i /defs/protos/events \
		-d /defs/protos/services \
		-l descriptor_set \
		-o /defs/build/go/protoset/ \
		--descr-include-imports \
		--descr-include-source-info \
		--descr-filename events.protoset \
		protos/**/*.proto

.PHONY: local
local: description = Compile protobuf schemas for Go and copy to plumber/grpc-collector projects
local: generate/go
local:
	cp -R build/go/protos/ ~/Code/plumber/vendor/github.com/batchcorp/collector-schemas/build/go/protos/
	cp -R build/go/protos/ ~/Code/grpc-collector/vendor/github.com/batchcorp/collector-schemas/build/go/protos/

.PHONY: clean-go
clean-go: description = Remove all go build artifacts
clean-go:
	rm -rf ./build/go/*

.PHONY: lint
lint: description = Run protolint
lint:
	protolint protos/
