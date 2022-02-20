package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	myApi "go-performance/internal/grpc/hello/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"net/http"
)

type HelloServiceImpl struct {
	*Authentication
}

func (h *HelloServiceImpl) Hello(ctx context.Context, in *myApi.String) (*myApi.String, error) {

	if err := h.Auth(ctx); err != nil {
		return nil, err
	}

	reply := &myApi.String{
		Value: "hello: " + in.GetValue(),
	}
	return reply, nil
}

func (h *HelloServiceImpl) Channel(stream myApi.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &myApi.String{Value: args.GetValue() + " is from stream reply"}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {

	grpcServer := grpc.NewServer()
	myApi.RegisterHelloServiceServer(grpcServer, &HelloServiceImpl{
		Authentication: &Authentication{
			User:     "sharpe",
			Password: "123456",
		},
	})

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		_ = grpcServer.Serve(lis)
	}()

	// 启动http 服务
	gwMux := runtime.NewServeMux(DefaultHTTPServeMuxOpt()...)
	err = myApi.RegisterHelloServiceHandlerFromEndpoint(context.Background(), gwMux, "127.0.0.1:1234",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", gwMux))

}

func DefaultHTTPServeMuxOpt() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		runtime.WithOutgoingHeaderMatcher(func(s string) (string, bool) {
			return s, true
		}),
		runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
			return s, true
		}),
	}
}

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

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var appid, appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}

	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != a.User || appkey != a.Password {
		return status.Errorf(codes.Unauthenticated, "invalid user")
	}
	return nil
}
