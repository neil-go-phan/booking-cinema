package routes

import (
	"booking-cinema-backend/api/handler"

	"github.com/gin-gonic/gin"
)

type RoleRoutes struct {
	roleHandler handler.RoleHandler
}

func NewRoleRoutes(roleHandler *handler.RoleHandler) *RoleRoutes{
	roleRoutes := &RoleRoutes{
		roleHandler: *roleHandler,
	}
	return roleRoutes
}

func (roleRoute *RoleRoutes)Setup(r *gin.Engine) {
	roleRoutes := r.Group("role")
	{
		roleRoutes.GET("roles", roleRoute.roleHandler.GetRoles)
	}
}