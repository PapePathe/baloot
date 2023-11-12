package main

import (
	"fmt"

	"pathe.co/zinx/zrpc"
)

func main() {
	addrs := []string{"localhost:50051", "localhost:50052"}
	cli, _ := zrpc.NewZGrpcClient("gametake", "baloot.GameTakeLearning", addrs)

	for i := 0; i < 1; i++ {
		err := cli.Send()

		fmt.Println(err)
	}

	fmt.Println(cli)
}
