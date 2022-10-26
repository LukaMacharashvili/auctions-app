package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/LukaMacharashvili/admin/pkg/config"
	"github.com/LukaMacharashvili/admin/pkg/pb"
	"github.com/LukaMacharashvili/admin/pkg/repositories"
	"github.com/LukaMacharashvili/admin/pkg/services"
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

	ar := repositories.AdminRepository{
		DB: db,
	}
	s := services.Server{
		AdminRepo: &ar,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAdminServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
