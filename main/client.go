package main

import (
	"context"
	"github.com/SDur/score-keeper/pb"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"time"

	// "google.golang.org/grpc/keepalive"
	"google.golang.org/grpc"
)

func runClient(servers []string) {
	target := servers[0]

	// use the simple LB
	if len(servers) > 1 {
		//target = simple.Target(servers)
	}
	conn, err := grpc.Dial(
		target,

		// some options
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),

		// block until connected
		grpc.WithBlock(),
	)
	if err != nil {
		log.Panicf("dial err: %s", err)
	}
	defer conn.Close()
	go printStateChange(conn, "conn")

	//client := pb.NewEchoClient(conn)
	client := pb.NewScoreKeeperServiceClient(conn)

	err = postScores(client)

	p1 := pb.Person{
		Firstname: "Dom",
		Lastname:  "Kop",
	}
	getScore(client, p1)

	p2 := pb.Person{
		Firstname: "Kaas",
		Lastname:  "Kop",
	}
	getScore(client, p2)
}

func getScore(client pb.ScoreKeeperServiceClient, p pb.Person) {
	person, err := client.GetScore(context.Background(), &p)
	if err != nil {
		log.Printf("error: %s", err)
	}
	log.Printf("Getting score for player [%s]", p.Firstname)
	log.Printf("Received score: [%d], for player: [%s]", person.GetWins(), person)
	time.Sleep(time.Second)
}

func postScores(client pb.ScoreKeeperServiceClient) error {
	log.Printf("---")
	log.Printf("Storing score")
	_, err := client.StoreScore(context.Background(), getMatchResult(5, 10))
	log.Printf("Storing score")
	_, err = client.StoreScore(context.Background(), getMatchResult(10, 7))
	log.Printf("Storing score")
	_, err = client.StoreScore(context.Background(), getMatchResult(10, 6))
	log.Printf("Storing score")
	_, err = client.StoreScore(context.Background(), getMatchResult(10, 4))
	if err != nil {
		log.Printf("error: %s", err)
		time.Sleep(time.Second * 5)
	}
	time.Sleep(time.Second)
	return err
}

func getMatchResult(rt1 int32, rt2 int32) *pb.MatchResult {
	p1 := pb.Person{
		Firstname: "Kaas",
		Lastname:  "Kop",
	}
	p2 := pb.Person{
		Firstname: "Stom",
		Lastname:  "Kop",
	}
	p3 := pb.Person{
		Firstname: "Dom",
		Lastname:  "Kop",
	}
	p4 := pb.Person{
		Firstname: "Pannenkoek",
		Lastname:  "Kop",
	}
	t1 := pb.MatchResult_Team{
		Persons: []*pb.Person{&p1, &p2},
		Score:   rt1,
	}
	t2 := pb.MatchResult_Team{
		Persons: []*pb.Person{&p3, &p4},
		Score:   rt2,
	}
	r := pb.MatchResult{
		Teams: []*pb.MatchResult_Team{&t1, &t2},
	}
	return &r
}

func printStateChange(conn *grpc.ClientConn, name string) {
	log.Printf("conn stat: %s", conn.GetState())
	for {
		state := conn.GetState()
		if conn.WaitForStateChange(context.Background(), state) {
			log.Printf("[%s] stage change %s->%s",
				name, state, conn.GetState())
		}
	}
}
