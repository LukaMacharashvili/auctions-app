package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/LukaMacharashvili/auth/pkg/config"
	"github.com/LukaMacharashvili/auth/pkg/helpers"
	"github.com/LukaMacharashvili/auth/pkg/pb"
	"github.com/LukaMacharashvili/auth/pkg/repositories"
	"github.com/LukaMacharashvili/auth/pkg/services"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	jwt := helpers.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	db, err := sql.Open("mysql", c.DBUrl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connected to DB")

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	ar := repositories.AuthRepository{
		DB: db,
	}
	s := services.Server{
		AuthRepo: &ar,
		Jwt:      &jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
