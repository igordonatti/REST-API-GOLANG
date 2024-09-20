package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igordonatti/REST-API/api/user/dto"
)

type UserController struct {
	Service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) MountRoutes(group *gin.Engine) {
	group.POST("/user", c.CreateUser)
	group.GET("/users", c.GetAllUsers)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var request dto.CreateUserDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.Service.CreateUser(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.Service.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
