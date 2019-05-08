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
	log.Println("Receiving request for storing a score")
	t1 := req.Teams[0]
	t2 := req.Teams[1]

	if t1.Score > t2.Score {
		log.Println("Team 1 has won")
		cache[t1.Persons[0].Firstname] = cache[t1.Persons[0].Firstname] + 1
		cache[t1.Persons[1].Firstname] = cache[t1.Persons[1].Firstname] + 1
	} else {
		log.Println("Team 2 has won")
		cache[t2.Persons[0].Firstname] = cache[t2.Persons[0].Firstname] + 1
		cache[t2.Persons[1].Firstname] = cache[t2.Persons[1].Firstname] + 1
	}
	return new(pb.Empty), nil
}

func (*service) GetScore(ctx context.Context, req *pb.Person) (*pb.Person, error) {
	log.Println("Receiving request to get score ")
	p := pb.Person{
		Firstname: "Sjaak",
		Lastname:  "Choco",
		Wins:      3,
	}
	result, ok := cache[req.Firstname]
	if !ok {
		return &p, nil
	}
	req.Wins = result
	log.Printf("Player %s has won %s times", req.Firstname, result)
	return req, nil
}

var cache = make(map[string]int32)

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
