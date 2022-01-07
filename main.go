package main

import "C"

import (
  "context"
  "flag"
  "net/http"

  "github.com/golang/glog"
  "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

)

var (
  // command-line options:
  // gRPC server endpoint
  grpcServerEndpoint = flag.String("grpc-server-endpoint",  "0.0.0.0:50051", "gRPC server endpoint")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  // Register gRPC server endpoint
  // Note: Make sure the gRPC server is running properly and accessible
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
  err := RegisterPathServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
  if err != nil {
    return err
  }

  // Start HTTP server (and proxy calls to gRPC server endpoint)
  return http.ListenAndServe(":8081", mux)
}

type server struct {
	ch     <-chan error
	cancel func()
}

var (
	servers = make(map[int]*server)
)

//export SpawnGrpcGateway
func SpawnGrpcGateway() int {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan error, 1)
	go func(ch chan<- error) {
		defer close(ch)
		if err := run(); err != nil {
			glog.Error("grpc-gateway failed with an error: %v", err)
			ch <- err
		}
	}(ch)

	handle := len(servers) + 1
	servers[handle] = &server{
		ch:     ch,
		cancel: cancel,
	}
	return handle
}

//export WaitForGrpcGateway
func WaitForGrpcGateway(handle int) bool {
	s, ok := servers[handle]
	if !ok {
		glog.Errorf("invalid handle: %d", handle)
		return false
	}
	s.cancel()
	err := <-s.ch
	return err == nil
}

func main() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}
