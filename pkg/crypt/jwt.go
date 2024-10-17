package crypt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/kitae0522/online-library-be/pkg/utils"
)

const JWTExpirationInSec = 60 * 60 * 24 * 7

func NewToken(userRole, userUUID string, secretKey []byte) (string, error) {
	expiration := time.Duration(JWTExpirationInSec) * time.Second
	claims := jwt.MapClaims{
		"uuid":      userUUID,
		"expiredAt": time.Now().Add(expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseJWT(jwtToken string) (string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ErrUnexpectedSigningMethod
		}
		return []byte("tempSecret"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if uuid, ok := claims["uuid"].(string); ok {
			return uuid, nil
		}
	}
	return "", utils.ErrInvalidTokenClaims
}
