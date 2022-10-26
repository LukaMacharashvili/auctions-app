package admin

import (
	"fmt"

	"github.com/LukaMacharashvili/api-gateway/pkg/admin/pb"
	"github.com/LukaMacharashvili/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AdminServiceClient
}

func InitServiceClient(c *config.Config) pb.AdminServiceClient {
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAdminServiceClient(cc)
}
