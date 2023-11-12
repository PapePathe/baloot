package zrpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gt_proto "pathe.co/zinx/proto/gametake/v1"
)

type ZGrpcClient struct {
	conn *grpc.ClientConn
}

func NewZGrpcClient(scheme, serviceName string, addrs []string) (*ZGrpcClient, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s://%s", scheme, serviceName),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	fmt.Println(conn)

	if err != nil {
		return nil, err
	}

	return &ZGrpcClient{conn: conn}, nil
}

func (c ZGrpcClient) Send() error {
	cli := gt_proto.NewGameTakeLearningClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	resp, err := cli.RecommendGameTake(ctx, &gt_proto.RecommendGameTakeRequest{
		Cards: []*gt_proto.Card{
			&gt_proto.Card{Type: "V", Color: "Pique"},
			&gt_proto.Card{Type: "V", Color: "Pique"},
			&gt_proto.Card{Type: "V", Color: "Pique"},
			&gt_proto.Card{Type: "V", Color: "Pique"},
			&gt_proto.Card{Type: "V", Color: "Pique"},
		},
	})

	if err != nil {
		return err
	}

	fmt.Println(resp)
	resp, err = cli.RecommendGameTake(ctx, &gt_proto.RecommendGameTakeRequest{})
	fmt.Println(resp)

	return nil
}
