package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/LukaMacharashvili/api-gateway/pkg/wallet/pb"
	"github.com/gin-gonic/gin"
)

type CreateRequestBody struct {
	Amount float64 `json:"amount"`
}

func Create(ctx *gin.Context, c pb.WalletServiceClient) {
	b := CreateRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("user id not found"))
		return
	}

	res, err := c.Create(context.Background(), &pb.CreateRequest{
		RaiseReq: &pb.RaiseReq{
			Amount:  b.Amount,
			OwnerId: userId.(int64),
			Status:  "Pending",
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
