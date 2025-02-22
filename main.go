package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	pb "gitlab.rnd.gic.ericsson.se/go-grpc-gateway/server"
	"google.golang.org/grpc"
)

var (
	serverHost  = flag.String("addr", "0.0.0.0", "The server host")
	proxyPort   = flag.String("proxyPort", "50051", "The grpc proxy server port")
	backendPort = flag.String("backendPort", "50052", "The grpc backend server port")
	client      = flag.Bool("client", false, "Run client")
)

type proxyServer struct {
	pb.UnimplementedProxyPingServer
}

func (s *proxyServer) SayPing(_ context.Context, msg *pb.PingMessage) (*pb.PingMessageReply, error) {
	log.Printf("Proxy received message: %s", msg.Id)
	select {
	// channel is not full
	case proxyChan <- msg:
		return &pb.PingMessageReply{Ok: true}, nil
	default:
		go writeToFile(msg)
		return &pb.PingMessageReply{Ok: true}, nil
	}
}

type backendServer struct {
	pb.UnimplementedBackendPingServer
}

func (s *backendServer) SayPing(_ context.Context, msg *pb.PingMessage) (*pb.PingMessageReply, error) {
	log.Printf("Backend received message: %s", msg.Id)
	return &pb.PingMessageReply{Ok: true}, nil
}

func writeToFile(msg *pb.PingMessage) {
	log.Printf("Writing message to file: %s", msg.Id)
}

func sendToGrpcBackendFromDisk() {
	select {}
}

func sendToGrpcBackend() {
	client, err := pingClient(backendPort)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case msg, ok := <-proxyChan:
			if ok {
				reply, err := sendBackendPing(client, msg)
				if err != nil || !reply.Ok {
					writeToFile(msg)
				}
			} else {
				fmt.Println("Channel is closed!")
			}
		}
	}
}

const proxyChanSize = 2

var proxyChan chan *pb.PingMessage

func grpcProxyServer() {
	listen, err := net.Listen("tcp", *serverHost+":"+*proxyPort)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterProxyPingServer(s, &proxyServer{})
	log.Fatal(s.Serve(listen))
}
func grpcBackendServer() {
	listen, err := net.Listen("tcp", *serverHost+":"+*backendPort)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterBackendPingServer(s, &backendServer{})
	log.Fatal(s.Serve(listen))
}

func main() {

	flag.Parse()

	if *client {
		proxyClient, err := pingClient(proxyPort)
		if err != nil {
			panic(err)
		}
		sendProxyPing(proxyClient, &pb.PingMessage{Id: generateId(), Body: "Hello!"})
	} else {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}

		log.Printf("Starting grpc backend on %s:%s", *serverHost, *backendPort)
		go grpcBackendServer()

		log.Printf("Starting backend channel sender")
		go sendToGrpcBackend()

		log.Printf("Starting backend file sender")
		go sendToGrpcBackendFromDisk()

		log.Printf("Starting grpc proxy on %s:%s", *serverHost, *proxyPort)
		proxyChan = make(chan *pb.PingMessage, proxyChanSize)
		go grpcProxyServer()

		select {}
	}
}
