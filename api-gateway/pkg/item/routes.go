package item

import (
	"github.com/LukaMacharashvili/api-gateway/pkg/auth"
	"github.com/LukaMacharashvili/api-gateway/pkg/config"
	"github.com/LukaMacharashvili/api-gateway/pkg/item/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/item")
	routes.Use(a.AuthRequired)
	routes.GET("/", svc.LoadItems)
	routes.GET("/:id", svc.FetchItem)
	routes.POST("/", svc.CreateItem)
	routes.DELETE("/:id", svc.DeleteItem)
	routes.PUT("/bid/:id", svc.BidItem)
	routes.PUT("/:id", svc.FinishAuction)
}

func (svc *ServiceClient) LoadItems(ctx *gin.Context) {
	routes.Load(ctx, svc.Client)
}

func (svc *ServiceClient) FetchItem(ctx *gin.Context) {
	routes.Fetch(ctx, svc.Client)
}

func (svc *ServiceClient) CreateItem(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteItem(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}

func (svc *ServiceClient) BidItem(ctx *gin.Context) {
	routes.Bid(ctx, svc.Client)
}

func (svc *ServiceClient) FinishAuction(ctx *gin.Context) {
	routes.FinishAuction(ctx, svc.Client)
}
