DOCKER := $(shell which docker)

.PHONY: build

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=0.11.4
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

build:
	GOARCH=arm64 GOOS=linux go build -o ./build/accountsd ./cmd/accountsd

docker: build
	docker container rm "localnet" -f
	docker build -t accounts:test .
	docker run -d --name "localnet" -t accounts:test
	docker exec -it "localnet" sh

exp: build
	docker container rm "localnet"
	docker build --no-cache -t accounts:test .
	docker run -it --name "localnet" -t accounts:test