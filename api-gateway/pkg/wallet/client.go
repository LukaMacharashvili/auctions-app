package wallet

import (
	"fmt"

	"github.com/LukaMacharashvili/api-gateway/pkg/config"
	"github.com/LukaMacharashvili/api-gateway/pkg/wallet/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.WalletServiceClient
}

func InitServiceClient(c *config.Config) pb.WalletServiceClient {
	cc, err := grpc.Dial(c.WalletSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewWalletServiceClient(cc)
}
