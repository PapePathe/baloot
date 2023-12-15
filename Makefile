generate-proto:
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative   --go_out=paths=source_relative:. proto/gametake_history/v1/gametakehistory.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative   --go_out=paths=source_relative:. proto/gametake/v1/gametake.proto

start:
	ZINX_STATIC_PATH=./public ZINX_TAKES_SERVER=localhost:50052 go run cmd/websocket/server.go
