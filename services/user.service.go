package services

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	*db.Queries
	*models.JWTHandler
}

// constructor
func NewUserService(dbConn *pgxpool.Conn, jwt *models.JWTHandler) *UserService {
	return &UserService{
		Queries:    db.New(dbConn),
		JWTHandler: jwt,
	}
}

func (us *UserService) Signup(userReq models.CreateUserReq) (*models.UserResponse, *models.Error) {
	//1. check is user already registered
	foundUser, _ := us.FindUserByUsername(context.Background(), &userReq.UserName)

	//2. already registered
	if foundUser.UserID != 0 {
		return &models.UserResponse{}, models.NewError(models.ErrDataNotFound)
	}

	//3. belum signup
	args := &db.CreateUserParams{
		UserName:     &userReq.UserName,
		UserPassword: userReq.UserPassword,
	}

	newUser, err := us.CreateUser(context.Background(), *args)
	if err != nil {
		return &models.UserResponse{}, models.NewError(err)
	}

	response := &models.UserResponse{
		UserID:    newUser.UserID,
		UserName:  newUser.UserName,
		UserPhone: newUser.UserPhone,
	}

	return response, nil

}

func (us UserService) Signin(userReq models.CreateUserReq) (*models.UserResponse,
	*models.Error) {

	args := &db.FindUserByUserPasswordParams{
		UserName:     &userReq.UserName,
		UserPassword: userReq.UserPassword,
	}

	foundUser, _ := us.FindUserByUserPassword(context.Background(), *args)

	if foundUser.UserID == 0 {
		return &models.UserResponse{}, models.NewError(models.ErrInvalidUserPassword)
	}

	accessToken, error := us.JWTHandler.GenerateJWT(*foundUser.UserName)

	if error != nil {
		return &models.UserResponse{},
			models.NewError(models.ErrFailedGenerateToken)
	}

	argsUpdateToken := &db.UpdateTokenParams{
		UserID:    foundUser.UserID,
		UserToken: &accessToken,
	}

	//update to table users, update kolom user_token
	_, err := us.UpdateToken(context.Background(), *argsUpdateToken)

	if err != nil {
		return &models.UserResponse{},
			models.NewError(models.ErrFailedGenerateToken)
	}

	response := &models.UserResponse{
		UserID:       foundUser.UserID,
		UserName:     foundUser.UserName,
		UserPassword: foundUser.UserPassword,
		UserPhone:    foundUser.UserPhone,
		UserToken:    &accessToken,
	}

	return response, nil
}

func (us UserService) Logout(accessToken string) *models.Error {

	err := us.DeleteToken(context.Background(), &accessToken)
	if err != nil {
		return models.NewError(models.ErrUpdateToken)
	}

	return nil
}

func (us UserService) GetProfileService(c *gin.Context) (*models.UserResponse, *models.Error) {
	var userID string
	token := us.JWTHandler.GetJWTFromHeader(c)
	if token != "" {
		userID = us.JWTHandler.GetIDFromToken(token)
	}

	foundUser, err := us.FindUserByUsername(context.Background(), &userID)
	if err != nil {
		return &models.UserResponse{}, models.NewError(err)
	}

	result := &models.UserResponse{
		UserID:    foundUser.UserID,
		UserName:  foundUser.UserName,
		UserPhone: foundUser.UserPhone,
	}

	return result, nil
}
