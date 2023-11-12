package zrpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pathe.co/zinx/pkg/cards"
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

	if err != nil {
		return nil, err
	}

	return &ZGrpcClient{conn: conn}, nil
}

func (c ZGrpcClient) Send() error {
	cli := gt_proto.NewGameTakeLearningClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	set := cards.CardSet{Cards: cards.CardSet{}.Distribute()}
	hand := set.Serve()
	resp, err := cli.RecommendGameTake(ctx, &gt_proto.RecommendGameTakeRequest{
		Cards: []*gt_proto.Card{
			&gt_proto.Card{Type: hand[0].Genre, Color: hand[0].Couleur},
			&gt_proto.Card{Type: hand[1].Genre, Color: hand[1].Couleur},
			&gt_proto.Card{Type: hand[2].Genre, Color: hand[2].Couleur},
			&gt_proto.Card{Type: hand[3].Genre, Color: hand[3].Couleur},
			&gt_proto.Card{Type: hand[4].Genre, Color: hand[4].Couleur},
		},
	})

	if err != nil {
		return err
	}

	fmt.Println(hand)
	for _, recommandation := range resp.AvailableTakes {
		fmt.Println(recommandation)
	}

	return nil
}
