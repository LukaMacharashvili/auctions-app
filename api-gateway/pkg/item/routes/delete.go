package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/LukaMacharashvili/api-gateway/pkg/item/pb"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context, c pb.ItemServiceClient) {
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

	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		OwnerId: userId.(int64),
		Id:      id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
