package grpc

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/common/productpb"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/config"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-commodity-mservice/server/log"
	"google.golang.org/grpc"
)

func Init(exitSig chan os.Signal) {
	address := fmt.Sprintf("%s:%d", config.Config.GrpcConfig.Host, config.Config.GrpcConfig.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Logger.Fatalf("Failed to listen: %v", err)
		exitSig <- os.Interrupt
		return
	}
	// Set up gRPC options for timeout and connection pooling
	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Duration(config.Config.GrpcConfig.ConnectTimeout) * time.Second), // Set a connection timeout
		grpc.MaxConcurrentStreams(uint32(config.Config.GrpcConfig.MaxPoolSize)),                      // Set maximum concurrent streams
		grpc.MaxRecvMsgSize(1024 * 1024), // Set maximum receive message size (1MB here)
		grpc.MaxSendMsgSize(1024 * 1024), // Set maximum send message size (1MB here)
	}
	grpcServer := grpc.NewServer(opts...)
	productpb.RegisterProductServiceServer(grpcServer, &ProductService{})

	log.Logger.Infof("Product RPC Server is running on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Logger.Fatal("Failed to serve: %v", err)
		exitSig <- os.Interrupt
	}
}
