package main

import (
	"database/sql"
	"flag"
	"git.neds.sh/matty/entain/sporting/db"
	"git.neds.sh/matty/entain/sporting/proto/sporting"
	"git.neds.sh/matty/entain/sporting/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net" 
)

var (
	grpcEndpoint = flag.String("grpc-endpoint", "localhost:10000", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("failed running grpc server: %s", err)
	}
}

func run() error {
	conn, err := net.Listen("tcp", ":10000")
	if err != nil {
		return err
	}

	sportingDB, err := sql.Open("sqlite3", "./db/sporting.db")
	if err != nil {
		return err
	}

	sportsRepo := db.NewSportsRepo(sportingDB)
	if err := sportsRepo.Init(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	sporting.RegisterSportingServer(
		grpcServer,
		service.NewSportingService(
			sportsRepo,
		),
	)

	log.Infof("gRPC server listening on: %s", *grpcEndpoint)

	if err := grpcServer.Serve(conn); err != nil {
		return err
	}

    return nil
}
