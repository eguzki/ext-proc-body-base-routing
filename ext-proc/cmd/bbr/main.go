package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	extProcPb "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/gateway-api-inference-extension/pkg/bbr/handlers"
)

func initLogging() {
	logger := zap.New(zap.UseDevMode(true))
	ctrl.SetLogger(logger)
}

func main() {
	initLogging()
	logger := ctrl.Log
	// grpc server init
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	extProcPb.RegisterExternalProcessorServer(s, handlers.NewServer(false))

	//log.Println("Starting gRPC server on port :50051")
	logger.Info("Starting gRPC server on port :50051")

	// shutdown
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("caught sig: %+v", sig)
		log.Println("Wait for 1 second to finish processing")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	s.Serve(lis)
}
