package user_service

import (
	"errors"
	"gin-app/apis/v1/config/database"
	"gin-app/apis/v1/models"
	jwtUtil "gin-app/apis/v1/utils/jwt"

	"github.com/dgrijalva/jwt-go"

	"github.com/mitchellh/mapstructure"
)

type UserAuthToken struct {
	ID       int
	Username string
}

func GenerateUserAuthToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["username"] = user.Username
	return jwtUtil.CreateToken(claims)
}

func ValidateUserToken(tokenString string) (*models.User, error) {
	db := database.GetDB()
	decodedData, err := jwtUtil.DecodeToken(tokenString)
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
