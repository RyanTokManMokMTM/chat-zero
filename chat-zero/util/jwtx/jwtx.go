package jwtx

import "github.com/golang-jwt/jwt/v4"

func GenerateToken(expiredTime, iat int64, key string, payload map[string]any) (string, error) {
	claim := make(jwt.MapClaims)
	claim["iat"] = iat
	claim["exp"] = iat + expiredTime
	for v, k := range payload {
		claim[v] = k
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claim
	return token.SignedString([]byte(key))
}
