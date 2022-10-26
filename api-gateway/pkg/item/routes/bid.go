package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/LukaMacharashvili/api-gateway/pkg/item/pb"
	"github.com/gin-gonic/gin"
)

type BidRequestBody struct {
	Amount float64 `json:"amount"`
}

func Bid(ctx *gin.Context, c pb.ItemServiceClient) {
	b := BidRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("user id not found"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.Bid(context.Background(), &pb.BidRequest{
		Id:     id,
		Amount: b.Amount,
		UserId: userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
