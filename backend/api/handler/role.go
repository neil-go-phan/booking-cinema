package handler

import (
	"booking-cinema-backend/helper"
	roleservice "booking-cinema-backend/services/role"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	handler roleservice.RoleService
}

func NewRoleHandler(handler roleservice.RoleService) *RoleHandler {
	roleHandler := &RoleHandler{
		handler: handler,
	}
	return roleHandler
}	

func (roleHandler *RoleHandler) GetRoles(c *gin.Context) {
	roles, err := roleHandler.handler.ListRole()
	if err != nil {
		c.Error(errors.New(helper.ERROR_BAD_REQUEST.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, roles)
}