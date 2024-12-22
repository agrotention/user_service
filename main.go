package main

import (
	"log"
	"net"
	"os"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/service"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Constants
	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		log.Fatalln("error running server, DB_URI variable not set")
	}
	addr := os.Getenv("SERVICE_ADDRESS")
	if addr == "" {
		log.Fatalln("error running server, SERVICE_ADDRESS variable not set")
	}

	// Init database
	db, err := gorm.Open(postgres.Open(dbUri))
	if err != nil {
		log.Fatal(err.Error())
	}

	// Init listener
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Comp server
	server := service.NewServer(db)
	grpcServer := grpc.NewServer()
	user_proto.RegisterUserServiceServer(grpcServer, server)

	// Run server
	log.Printf("server running on %s", addr)
	log.Fatal(grpcServer.Serve(lis))
}
