// Main package
//
// Create a new grpc server that captures gametakes history
// and persist them to a storage system.
package main

import (
	"flag"
	"fmt"
	"log"

	proto "pathe.co/zinx/proto/gametake_history/v1"
	"pathe.co/zinx/zrpc"
)

var port = flag.Int("port", 50052, "The server port")

func main() {
	flag.Parse()

	rpcServer, err := zrpc.NewZGrpcServer([]int{*port})

	if err != nil {
		log.Fatalf(fmt.Sprintf("error creating rpc server %s", err))
	}

	for _, s := range rpcServer.Backends {
		proto.RegisterGameTakeHistoryServer(s, zrpc.NewGameTakeHistoryServer())
	}

	if err := rpcServer.Start(); err != nil {
		log.Fatalf(fmt.Sprintf("error starting rpc server %s", err))
	}
}
