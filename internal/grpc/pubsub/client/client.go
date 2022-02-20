package main

import (
	"context"
	"fmt"
	"go-performance/internal/grpc/pubsub/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"time"
)

func main() {

	go func() {
		time.Sleep(time.Second * 10)

		conn, err := grpc.Dial("127.0.0.1:1235", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		client := api.NewPubSubServiceClient(conn)
		_, err = client.Publish(context.Background(), &api.String{
			Value: "golang java c",
		})
		if err != nil {
			panic(err)
		}

		_, err = client.Publish(context.Background(), &api.String{
			Value: "docker k8s pod",
		})
		if err != nil {
			panic(err)
		}

		fmt.Println("Publish successfully")
	}()

	conn2, err := grpc.Dial("127.0.0.1:1235", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn2.Close()

	client2 := api.NewPubSubServiceClient(conn2)
	stream, err := client2.Subscribe(context.Background(), &api.String{Value: "golang"})
	if err != nil {
		panic(err)
	}

	for {
		reply, streamErr := stream.Recv()
		if streamErr != nil {
			if streamErr == io.EOF {
				fmt.Println("break")
				break
			}
			panic(streamErr)
		}
		fmt.Println("stream reply = ", reply.GetValue())
	}
}
