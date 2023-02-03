package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"simple_bank/api"
	db "simple_bank/db/sqlc"
	"simple_bank/gapi"
	"simple_bank/pb"
	"simple_bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load the config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewSQLStore(conn)
	//runGinServer(config, store)
	runGrpcServer(config, store)
}

func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	// this command look very simple, but it actually is pretty powerful,
	// because it allows the gRPC client to easily explore
	// what RPCs are available on the server, and how to call them, (like enable feature GraphQL Playground)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
