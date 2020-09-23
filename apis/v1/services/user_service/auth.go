package user_service

import (
	"errors"
	"gin-app/apis/v1/config/database"
	"gin-app/apis/v1/models"
	"gin-app/apis/v1/utils/jwt"

	"github.com/mitchellh/mapstructure"
)

type UserAuthToken struct {
	ID       int
	Username string
}

func ValidateUserToken(tokenString string) (*models.User, error) {
	db := database.GetDB()
	decodedData, err := jwt.DecodeToken(tokenString)
	if err != nil {
		return nil, err
	}

	var userAuthToken UserAuthToken
	dErr := mapstructure.Decode(decodedData, &userAuthToken)
	if dErr != nil {
		return nil, errors.New("Invalid Token Format")
	}

	u := models.User{}
	user, err := u.GetUser(db, userAuthToken.ID)
	if err != nil {
		return nil, errors.New("No User found with provided Authorization")
	}

	return user, nil
}
