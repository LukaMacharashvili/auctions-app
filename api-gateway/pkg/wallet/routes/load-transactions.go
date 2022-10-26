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

func LoadTransactions(ctx *gin.Context, c pb.WalletServiceClient) {
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

	res, err := c.LoadTransactions(context.Background(), &pb.LoadTransactionsRequest{
		OwnerId: ownerId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	var transactionList []*models.Transaction

	for {
		var transaction models.Transaction
		res, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Internal Server Error: %v", err)
		}

		transaction.Id = res.GetTransaction().GetId()
		transaction.Amount = res.GetTransaction().GetAmount()
		transaction.OwnerId = res.GetTransaction().GetOwnerId()

		transactionList = append(transactionList, &transaction)
	}

	ctx.JSON(http.StatusOK, &transactionList)
}
