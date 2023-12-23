CONFIG_PATH=${HOME}/.dumb_broker/

.PHONY: init
init:
	mkdir -p ${CONFIG_PATH}

.PHONY: gencert
gencert:
	cfssl gencert -initca test/ca-csr.json | cfssljson -bare ca
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=server \
		test/server-csr.json | cfssljson -bare server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=client \
		test/client-csr.json | cfssljson -bare client
	mv *.pem *.csr ${CONFIG_PATH}

.PHONY: test
test:
	go test -race ./...

generate-proto:
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=paths=source_relative:. proto/gametake_history/v1/gametakehistory.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=paths=source_relative:. proto/gametake/v1/gametake.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=paths=source_relative:. proto/dumb_broker/v1/log.proto

start:
	ZINX_STATIC_PATH=./public ZINX_TAKES_SERVER=localhost:50052 go run cmd/websocket/server.go
