package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/LukaMacharashvili/api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Set("Roles", res.Roles)

	ctx.Next()
}

type RoleMiddlewareConfig struct {
	role string
}

func InitRoleMiddleware(role string) RoleMiddlewareConfig {
	return RoleMiddlewareConfig{role}
}

func (c *RoleMiddlewareConfig) RoleCheck(ctx *gin.Context) {
	roles := ctx.GetString("Roles")
	rolesList := strings.Split(roles, " ")

	if contains(rolesList, "admin") {
		ctx.Set("isAdmin", true)
		ctx.Next()
	}

	if contains(rolesList, c.role) {
		ctx.Next()
	}

	ctx.AbortWithStatus(http.StatusForbidden)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
