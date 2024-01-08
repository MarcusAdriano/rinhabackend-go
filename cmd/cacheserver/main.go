package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"sync"

	pb "github.com/marcusadriano/rinhabackend-go/internal/cache"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCacheServiceServer
	cache map[string][]byte
	lock  sync.RWMutex
}

const (
	OkResult = "OK"
)

var (
	NotFoundErr = errors.New("not found")
)

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.cache[req.Key] = req.Value
	return &pb.PutResponse{
		Result: OkResult,
	}, nil
}
func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	s.lock.RLocker().Lock()
	defer s.lock.RLocker().Unlock()

	if value, ok := s.cache[req.Key]; ok {
		return &pb.GetResponse{
			Value: value,
		}, nil
	}
	return nil, NotFoundErr
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCacheServiceServer(s, &server{
		cache: make(map[string][]byte),
	})
	log.Info().Msgf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
