package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LukaMacharashvili/api-gateway/pkg/item/models"
	"github.com/LukaMacharashvili/api-gateway/pkg/item/pb"
	"github.com/gin-gonic/gin"
)

func Fetch(ctx *gin.Context, c pb.ItemServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.Fetch(context.Background(), &pb.FetchRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	var item models.Item
	item.Id = res.GetItem().GetId()
	item.Name = res.GetItem().GetName()
	item.Description = res.GetItem().GetDescription()
	item.StartingPrice = res.GetItem().GetStartingPrice()
	item.Image = res.GetItem().GetImage()
	item.Status = res.GetItem().GetStatus()
	item.HighestBid = res.GetItem().GetHighestBid()
	item.OwnerId = res.GetItem().GetOwnerId()
	item.HighestBidder = res.GetItem().GetHighestBidder()

	ctx.JSON(http.StatusOK, &item)
}
