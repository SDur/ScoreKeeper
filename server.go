package score_keeper

import (
	"context"
	"fmt"
	"github.com/SDur/score-keeper/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type server struct {
	port string
}

type service struct {
}

func (*service) StoreScore(ctx context.Context, req *pb.MatchResult) (*pb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreScore not implemented")
}

func (*service) GetScore(ctx context.Context, req *pb.Person) (*pb.Person, error) {
	p := pb.Person{
		Firstname: "Sjaak",
		Lastname:  "Choco",
		Wins:      3,
	}
	return &p, nil
}

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
