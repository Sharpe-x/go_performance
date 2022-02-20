package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/pkg/pubsub"
	"go-performance/internal/grpc/pubsub/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

type PubSubService struct {
	pub *pubsub.Publisher
}

func NewPubSubService() *PubSubService {
	return &PubSubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubSubService) Publish(ctx context.Context, arg *api.String) (*api.String, error) {
	fmt.Println(arg.GetValue())
	p.pub.Publish(arg.GetValue())

	return &api.String{}, nil
}

func (p *PubSubService) Subscribe(arg *api.String, stream api.PubSubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.Contains(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&api.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	api.RegisterPubSubServiceServer(grpcServer, NewPubSubService())

	lis, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatal(err)
	}
	_ = grpcServer.Serve(lis)
}
