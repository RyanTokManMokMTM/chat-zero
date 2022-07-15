package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/ryantokmanmokmtm/house-booking-service/common/ctxtool"
)

func TokenGenerate(userID, iat, exp int64, key string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + exp
	claims["iat"] = iat
	claims[ctxtool.JWTTokenUSerID] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}