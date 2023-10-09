pb:
	@protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto

build:
	@echo ">> Building GRPC..."
	@go build -o stockbit-grpc ./cmd/grpc
	@echo ">> Building API..."
	@go build -o stockbit-api ./cmd/api
