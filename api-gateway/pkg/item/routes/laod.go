package routes

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/LukaMacharashvili/api-gateway/pkg/item/models"
	"github.com/LukaMacharashvili/api-gateway/pkg/item/pb"
	"github.com/gin-gonic/gin"
)

func Load(ctx *gin.Context, c pb.ItemServiceClient) {
	res, err := c.Load(context.Background(), &pb.LoadRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	var itemList []*models.Item

	for {
		var item models.Item
		res, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Internal Server Error: %v", err)
		}

		item.Id = res.GetItem().GetId()
		item.Name = res.GetItem().GetName()
		item.Description = res.GetItem().GetDescription()
		item.StartingPrice = res.GetItem().GetStartingPrice()
		item.Image = res.GetItem().GetImage()
		item.Status = res.GetItem().GetStatus()
		item.HighestBid = res.GetItem().GetHighestBid()
		item.OwnerId = res.GetItem().GetOwnerId()
		item.HighestBidder = res.GetItem().GetHighestBidder()

		itemList = append(itemList, &item)
	}

	ctx.JSON(http.StatusOK, &itemList)
}
