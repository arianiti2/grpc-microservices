MODULE_NAME = github.com/arianiti2/grpc-microservices


PROTO_DIR = api/v1
GEN_DIR = gen/go


PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc


GO := go


.PHONY: all proto build run clean

all: build


proto:
	@echo "Generating gRPC code from proto files..."
	@protoc \
		-I $(PROTO_DIR) \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto
	@echo "Proto generation complete."


build: proto
	@echo "Building gRPC server..."
	$(GO) build -o bin/server ./cmd/server
	@echo "Build complete. Binary: bin/server"


run: build
	@echo "Starting gRPC server..."
	./bin/server


clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -rf bin/*
	rm -rf $(GEN_DIR)/*
	@echo "Clean complete."
