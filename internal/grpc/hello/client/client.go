package main

import (
	"context"
	"fmt"
	myApi "go-performance/internal/grpc/hello/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"time"
)

// Authentication 实现PerRPCCredentials 接口 对每个grpc 方法进行认证
type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {

	/*	creds, err := credentials.NewClientTLSFromFile(
			"/Users/sharpe/workspace/dev_cloud/go_src/go_performance/internal/grpc"+
				"/hello/server/server.crt", "server.grpc.io")
		if err != nil {
			log.Fatal(err)
		}
	*/

	auth := &Authentication{
		User:     "sharpe",
		Password: "123456",
	}

	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(auth))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := myApi.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &myApi.String{Value: "hello"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply.GetValue())

	stream, err := client.Channel(context.Background())
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			if goErr := stream.Send(&myApi.String{Value: "hi"}); goErr != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		streamReply, streamErr := stream.Recv()
		if streamErr != nil {
			if streamErr == io.EOF {
				break
			}
			panic(streamErr)
		}
		fmt.Println(streamReply.GetValue())
	}

	select {}
}

//  openssl genrsa -out server.key 2048
//  openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" -key server.key -out server.crt

//  openssl genrsa -out client.key 2048
//  openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" -key client.key -out client.crt
