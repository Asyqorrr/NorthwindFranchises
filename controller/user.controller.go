package controller

import (
	"b30northwindapi/models"
	"b30northwindapi/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	serviceManager *services.ServiceManager
}

func NewUserController(serviceManager *services.ServiceManager) *UserController {
	return &UserController{
		serviceManager: serviceManager,
	}
}

func (uc *UserController) Signup(c *gin.Context) {
	var payload *models.CreateUserReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}
	args := &models.CreateUserReq{
		UserName:     payload.UserName,
		UserPassword: payload.UserPassword,
	}

	user, err := uc.serviceManager.Signup(*args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)

}

func (uc *UserController) Sigin(c *gin.Context) {
	var payload models.CreateUserReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &models.CreateUserReq{
		UserName:     payload.UserName,
		UserPassword: payload.UserPassword,
	}
	user, err := uc.serviceManager.Signin(*args)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Signout(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")

	err := uc.serviceManager.Logout(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (uc *UserController) GetProfile(c *gin.Context) {
	var userID string

	token := uc.serviceManager.JWTHandler.GetJWTFromHeader(c)
	if token != "" {
		userID = uc.serviceManager.GetIDFromToken(token)
	}

	foundUser, err := uc.serviceManager.UserService.FindUserByUsername(context.Background(), &userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	profile := &models.UserResponse{
		UserID:    foundUser.UserID,
		UserName:  foundUser.UserName,
		UserPhone: foundUser.UserPhone,
	}

	c.JSON(http.StatusOK, profile)
}
