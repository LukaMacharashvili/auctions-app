package routes

import (
	"context"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/LukaMacharashvili/api-gateway/pkg/wallet/models"
	"github.com/LukaMacharashvili/api-gateway/pkg/wallet/pb"
	"github.com/gin-gonic/gin"
)

func LoadRaiseReqs(ctx *gin.Context, c pb.WalletServiceClient) {
	ownerId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	userId, _ := ctx.Get("userId")

	isAdmin, _ := ctx.Get("isAdmin")

	if ownerId != userId.(int64) && !isAdmin.(bool) {
		ctx.AbortWithError(http.StatusForbidden, err)
		return
	}

	res, err := c.LoadRaiseReqs(context.Background(), &pb.LoadRaiseReqsRequest{
		OwnerId: ownerId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	var raiseReqList []*pb.RaiseReq
	for {
		var raiseReq models.RaiseReq
		res, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Internal Server Error: %v", err)
		}

		raiseReq.Id = res.GetRaiseReq().GetId()
		raiseReq.Amount = res.GetRaiseReq().GetAmount()
		raiseReq.OwnerId = res.GetRaiseReq().GetOwnerId()
		raiseReq.Status = res.GetRaiseReq().GetStatus()

		raiseReqList = append(raiseReqList, res.GetRaiseReq())
	}
	ctx.JSON(http.StatusOK, &raiseReqList)
}
