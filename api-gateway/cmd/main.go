package main

import (
	"log"

	"github.com/LukaMacharashvili/api-gateway/pkg/admin"
	"github.com/LukaMacharashvili/api-gateway/pkg/auth"
	"github.com/LukaMacharashvili/api-gateway/pkg/config"
	"github.com/LukaMacharashvili/api-gateway/pkg/item"
	"github.com/LukaMacharashvili/api-gateway/pkg/wallet"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, &c)
	item.RegisterRoutes(r, &c, authSvc)
	admin.RegisterRoutes(r, &c, authSvc)
	wallet.RegisterRoutes(r, &c, authSvc)

	r.Run(c.Port)
}
