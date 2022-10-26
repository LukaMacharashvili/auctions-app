package item

import (
	"fmt"

	"github.com/LukaMacharashvili/api-gateway/pkg/config"
	"github.com/LukaMacharashvili/api-gateway/pkg/item/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ItemServiceClient
}

func InitServiceClient(c *config.Config) pb.ItemServiceClient {
	cc, err := grpc.Dial(c.ItemSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewItemServiceClient(cc)
}
