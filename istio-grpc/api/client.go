package api

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "TODO: ADD APP ID"

// NewClient new grpc client
func NewClient(opts ...grpc.DialOption) (HelloWorldClient, error) {
	// client := warden.NewClient(cfg, opts...)
	// cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	// if err != nil {
	// 	return nil, err
	// }
	var conn *grpc.ClientConn
	var err error
	if conn, err = grpc.DialContext(context.TODO(), "192.168.81.46:18888", grpc.WithInsecure()); err != nil {
		fmt.Fprintf(os.Stderr, "warden client: dial %s error %v!", "grpc-server-svc:9312", err)
	}
	return NewHelloWorldClient(conn), nil
}

// 生成 gRPC 代码
//go:generate higo tool protoc --grpc --bm api.proto
