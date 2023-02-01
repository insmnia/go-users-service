package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/insmnia/go-users-service/config"
	"time"
)

type JWTClaim struct {
	Username string `json:"username"`
	UserId   string `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(username string, userId string, isRefresh bool) (tokenString string, err error) {
	var JwtConfig, jwtParseErr = config.LoadJWTConfig("./env")
	if jwtParseErr != nil {
		panic("Cannot load jwt config!")
	}
	var targetsTime = JwtConfig.AccessTokenLifetime
	if isRefresh {
		targetsTime = JwtConfig.RefreshTokenLifetime
	}

	expirationTime := time.Now().Add(time.Duration(targetsTime) * time.Hour)
	claims := &JWTClaim{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(JwtConfig.JwtSecret))
	return
}
func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	var JwtConfig, jwtParseErr = config.LoadJWTConfig(".")
	if jwtParseErr != nil {
		panic("Cannot load jwt config!")
	}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtConfig.JwtSecret), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
