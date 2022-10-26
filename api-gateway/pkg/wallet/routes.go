package wallet

import (
	"github.com/LukaMacharashvili/api-gateway/pkg/auth"
	"github.com/LukaMacharashvili/api-gateway/pkg/config"
	"github.com/LukaMacharashvili/api-gateway/pkg/wallet/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	roleMiddleware := auth.InitRoleMiddleware("admin")

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/wallet")
	routes.Use(a.AuthRequired)
	routes.GET("/transaction/:userId", svc.LoadTransactions)
	routes.GET("/raisereq/:userId", svc.LoadRaiseReqs)
	routes.POST("/raisereq", svc.Create)
	routes.PUT("/raisereq/:userId/:id/:status", roleMiddleware.RoleCheck, svc.ChangeInvoiceStatus)
}

func (svc *ServiceClient) LoadTransactions(ctx *gin.Context) {
	routes.LoadTransactions(ctx, svc.Client)
}

func (svc *ServiceClient) LoadRaiseReqs(ctx *gin.Context) {
	routes.LoadRaiseReqs(ctx, svc.Client)
}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}

func (svc *ServiceClient) ChangeInvoiceStatus(ctx *gin.Context) {
	routes.ChangeStatus(ctx, svc.Client)
}
