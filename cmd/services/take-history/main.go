package main

import (
	"flag"
	"fmt"
	"log"

	proto "pathe.co/zinx/proto/gametake_history/v1"
	"pathe.co/zinx/zrpc"
)

var port = flag.Int("port", 50051, "The server port")

func main() {
	flag.Parse()

	rpcServer, err := zrpc.NewZGrpcServer(*port)

	if err != nil {
		log.Fatalf(fmt.Sprintf("error creating rpc server %s", err))
	}

	proto.RegisterGameTakeHistoryServer(
		rpcServer.Server,
		zrpc.NewGameTakeHistoryServer(),
	)

	if err := rpcServer.Start(); err != nil {
		log.Fatalf(fmt.Sprintf("error starting rpc server %s", err))
	}
}
