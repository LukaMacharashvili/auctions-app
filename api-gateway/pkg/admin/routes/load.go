package routes

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/LukaMacharashvili/api-gateway/pkg/admin/models"
	"github.com/LukaMacharashvili/api-gateway/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func Load(ctx *gin.Context, c pb.AdminServiceClient) {
	res, err := c.Load(context.Background(), &pb.LoadRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	var users []*models.User
	for {
		var user models.User
		res, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}

		user.Id = res.GetUser().GetId()
		user.Username = res.GetUser().GetUsername()
		user.Email = res.GetUser().GetEmail()
		user.Password = res.GetUser().GetPassword()
		user.Roles = res.GetUser().GetRole()
		user.Balance = res.GetUser().GetBalance()

		users = append(users, &user)
	}

	ctx.JSON(http.StatusOK, &users)
}
