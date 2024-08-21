package services

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/middleware"
	"b30northwindapi/models"
	"context"
)

func (sm *StoreManager) Signup(userReq models.CreateUserReq) (*models.UserResponse, *models.Error) {
	//1. is user already signup
	foundUser, _ := models.Nullable(sm.FindUserByUsername(context.Background(), userReq.UserName))

	if foundUser.UserID != 0 {
		return &models.UserResponse{}, models.NewError(models.ErrUserAlreadyExist)
	}

	// user not exist
	argCreateUser := db.CreateUserParams{
		UserName:     userReq.UserName,
		UserPassword: userReq.UserPassword,
	}
	newUser, err := sm.CreateUser(context.Background(), argCreateUser)
	if err != nil {
		return &models.UserResponse{}, models.NewError(err)
	}

	response := &models.UserResponse{
		UserID:       newUser.UserID,
		UserPassword: newUser.UserPassword,
		UserPhone:    newUser.UserPhone,
	}
	return response, nil
}

func (sm *StoreManager) Signin(userReq models.CreateUserReq) (*models.UserResponse, *models.Error) {
	args := &db.FindUserByUserPasswordParams{
		UserName:     userReq.UserName,
		UserPassword: userReq.UserPassword,
	}

	foundUser, _ := sm.FindUserByUserPassword(context.Background(), *args)

	if foundUser.UserID == 0 {
		return &models.UserResponse{}, models.NewError(models.ErrInvalidUserPassword)
	}

	accessToken, error := middleware.GenerateJWT(*foundUser.UserName)

	if error != nil {
		return &models.UserResponse{},
			models.NewError(models.ErrFailedGenerateToken)
	}

	argsUpdateToken := &db.UpdateTokenParams{
		UserID:    foundUser.UserID,
		UserToken: &accessToken,
	}

	_, err := sm.UpdateToken(context.Background(), *argsUpdateToken)

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

func (sm *StoreManager) Signout(accessToken string) *models.Error {
	err := sm.DeleteToken(context.Background(), &accessToken)
	if err != nil {
		return models.NewError(models.ErrUpdateToken)
	}

	return nil
}

/* func generateAccessToken(username string) (string, *models.Error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(username), bcrypt.DefaultCost)
	if err != nil {
		return "", models.NewError(models.ErrFailedGenerateToken)
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}
*/
