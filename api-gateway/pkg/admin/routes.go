package admin

import (
	"github.com/LukaMacharashvili/api-gateway/pkg/admin/routes"
	"github.com/LukaMacharashvili/api-gateway/pkg/auth"
	"github.com/LukaMacharashvili/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	roleMiddlewareInit := auth.InitRoleMiddleware("admin")

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/admin")
	routes.Use(a.AuthRequired)
	routes.Use(roleMiddlewareInit.RoleCheck)
	routes.GET("/", svc.LoadUser)
	routes.POST("/", svc.CreateUser)
	routes.DELETE("/:id", svc.DeleteUser)
}

func (svc *ServiceClient) LoadUser(ctx *gin.Context) {
	routes.Load(ctx, svc.Client)
}

func (svc *ServiceClient) CreateUser(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}
