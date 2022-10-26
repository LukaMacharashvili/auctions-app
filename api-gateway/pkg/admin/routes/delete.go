package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LukaMacharashvili/api-gateway/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context, c pb.AdminServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
