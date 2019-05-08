package main

import (
	"context"
	"fmt"
	"github.com/SDur/score-keeper/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

type service struct {
}

func (*service) StoreScore(ctx context.Context, req *pb.MatchResult) (*pb.Empty, error) {
	t1 := req.Teams[0]
	t2 := req.Teams[1]

	if t1.Score > t2.Score {
		cache[t1.Persons[0]] = cache[t1.Persons[0]] + 1
		cache[t1.Persons[1]] = cache[t1.Persons[1]] + 1
	} else {
		cache[t2.Persons[0]] = cache[t2.Persons[0]] + 1
		cache[t2.Persons[1]] = cache[t2.Persons[1]] + 1
	}
	return new(pb.Empty), nil
}

func (*service) GetScore(ctx context.Context, req *pb.Person) (*pb.Person, error) {
	p := pb.Person{
		Firstname: "Sjaak",
		Lastname:  "Choco",
		Wins:      3,
	}
	result, ok := cache[req]
	if !ok {
		return &p, nil
	}
	req.Wins = result
	return req, nil
}

var cache = make(map[*pb.Person]int32)

func runServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Second * 10,
			Timeout:           time.Second * 20,
		}),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             time.Second,
				PermitWithoutStream: true,
			}),
		grpc.MaxConcurrentStreams(5),
	)

	pb.RegisterScoreKeeperServiceServer(s, new(service))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
