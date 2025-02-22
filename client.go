package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/norrbom/grpc-gateway/server"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func pingClient(port *string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.NewClient(*serverHost+":"+*port, opts...)
}

func generateId() string {
	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	return fmt.Sprintf("%d", r.Uint64())
}

func sendProxyPing(conn *grpc.ClientConn, msg *pb.PingMessage) (*pb.PingMessageReply, error) {
	client := pb.NewProxyPingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.SayPing(ctx, msg)
}

func sendBackendPing(conn *grpc.ClientConn, msg *pb.PingMessage) (*pb.PingMessageReply, error) {
	client := pb.NewBackendPingClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	time.Sleep(100 * time.Millisecond)
	return client.SayPing(ctx, msg)
}
