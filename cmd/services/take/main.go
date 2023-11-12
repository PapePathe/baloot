package main

import (
	"flag"
	"fmt"
	"log"

	proto "pathe.co/zinx/proto/gametake/v1"
	"pathe.co/zinx/zrpc"
)

var port = flag.Int("port", 50052, "The server port")

func main() {
	flag.Parse()

	rpcServer, err := zrpc.NewZGrpcServer([]int{50052, 50053, 50054, 50055})

	if err != nil {
		log.Fatalf(fmt.Sprintf("error creating rpc server %s", err))
	}

	for _, s := range rpcServer.Backends {
		proto.RegisterGameTakeLearningServer(
			s,
			zrpc.NewRecommendGameTakeServer(),
		)
	}

	if err := rpcServer.Start(); err != nil {
		log.Fatalf(fmt.Sprintf("error starting rpc server %s", err))
	}

}
