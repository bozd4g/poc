package usercontroller

import (
	"github.com/bozd4g/poc/testcontainer/cmd/testcontainers/internal/application/userservice"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Init(e *gin.Engine)
}

type UserController struct {
	service userservice.IUserService
}
