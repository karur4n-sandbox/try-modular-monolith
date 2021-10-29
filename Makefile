GRPC_OUT_DIR := ./internal/pkg

.PHONY: protoc
protoc:
	protoc \
		--go_out=$(GRPC_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GRPC_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		./proto/seminar/v1/*.proto
