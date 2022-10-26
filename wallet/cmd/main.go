package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/LukaMacharashvili/wallet/pkg/config"
	"github.com/LukaMacharashvili/wallet/pkg/pb"
	"github.com/LukaMacharashvili/wallet/pkg/repositories"
	"github.com/LukaMacharashvili/wallet/pkg/services"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
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

	rr := repositories.RaiseReqRepository{
		DB: db,
	}

	tr := repositories.TransactionRepository{
		DB: db,
	}

	ur := repositories.UserRepository{
		DB: db,
	}

	s := services.Server{
		TransactionRepo: &tr,
		RaiseReqRepo:    &rr,
		UserRepo:        &ur,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterWalletServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
