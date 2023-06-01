package utils

import (
	"errors"
	"golang-store/model/entity"
	"golang-store/model/web"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(request entity.User, expired time.Duration) (*string, error) {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	var APPLICATION_NAME = "GOLANG-STORE"

	expiredTime := time.Now().Add(time.Hour * expired).Unix()
	claims := &web.TokenClaims{
		Id:        int16(request.Id),
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: expiredTime,
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokesStr, err := tokens.SignedString(jwtTokenSecret)
	if err != nil {
		return nil, err
	}
	return &tokesStr, nil

}

func TokenClaims(userToken string) (*web.TokenClaims, error) {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	claims := &web.TokenClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtTokenSecret, nil
	})
	if err != nil {
		return nil, errors.New("failed claims tokens")
	}

	if !token.Valid {
		return nil, errors.New("tokens not valid")
	}

	return claims, nil
}
