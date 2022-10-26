package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/LukaMacharashvili/api-gateway/pkg/item/pb"
	"github.com/gin-gonic/gin"
)

type CreateRequestBody struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	StartingPrice float64 `json:"startingPrice"`
	Image         string  `json:"image"`
}

func Create(ctx *gin.Context, c pb.ItemServiceClient) {
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
		Item: &pb.Item{
			Name:          b.Name,
			Description:   b.Description,
			StartingPrice: b.StartingPrice,
			Image:         b.Image,
			OwnerId:       userId.(int64),
			HighestBidder: userId.(int64),
		},
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
