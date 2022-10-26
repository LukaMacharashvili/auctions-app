package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LukaMacharashvili/api-gateway/pkg/wallet/pb"
	"github.com/gin-gonic/gin"
)

func ChangeStatus(ctx *gin.Context, c pb.WalletServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	status := ctx.Param("status")

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.ChangeStatus(context.Background(), &pb.ChangeStatusRequest{
		Id:     id,
		Status: status,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
