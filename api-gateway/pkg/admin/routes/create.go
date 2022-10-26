package routes

import (
	"context"
	"net/http"

	"github.com/LukaMacharashvili/api-gateway/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

type CreateRequestBody struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

func Create(ctx *gin.Context, c pb.AdminServiceClient) {
	b := CreateRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Create(context.Background(), &pb.CreateRequest{
		User: &pb.User{
			Username: b.Username,
			Email:    b.Email,
			Password: b.Password,
			Role:     "admin",
			Balance:  b.Balance,
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
