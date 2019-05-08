package score_keeper

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

		// backoff policy
		// grpc.WithBackoffConfig(grpc.BackoffConfig{
		// 	MaxDelay: time.Second,
		// }),
		// grpc.WithBackoffMaxDelay(time.Second),

		// disable healthcheck, seems not working
		// grpc.WithDisableHealthCheck(),

		// maybe works under high corrency
		// grpc.WithDisableRetry(),

		// care of this config, read the comments carefully
		// grpc.WithKeepaliveParams(keepalive.ClientParameters{
		// 	Time:    time.Second,
		// 	Timeout: time.Second * 5,
		// }),
	)
	if err != nil {
		log.Panicf("dial err: %s", err)
	}
	defer conn.Close()
	go printStateChange(conn, "conn")

	//client := pb.NewEchoClient(conn)
	client := pb.NewScoreKeeperServiceClient(conn)
	log.Printf("---")
	_, err = client.StoreScore(context.Background(), getMatchResult())
	if err != nil {
		log.Printf("error: %s", err)
		time.Sleep(time.Second * 5)
	}
	time.Sleep(time.Second)
	p1 := pb.Person{
		Firstname: "Kaas",
		Lastname:  "Kop",
	}
	person, err := client.GetScore(context.Background(), &p1)
	if err != nil {
		log.Printf("error: %s", err)
		time.Sleep(time.Second * 5)
	}
	log.Printf("Received score: %s, for player: %s", person.GetWins(), person)
}

func getMatchResult() *pb.MatchResult {
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
		Score:   5,
	}
	t2 := pb.MatchResult_Team{
		Persons: []*pb.Person{&p3, &p4},
		Score:   10,
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
