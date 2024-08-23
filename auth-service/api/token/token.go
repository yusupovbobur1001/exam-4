package token

import (
	"auth_service/config"
	"auth_service/model"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.StandardClaims
}

func GenerateJWT(user *model.LoginResponse) *model.Tokens {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["first_name"] = user.FirstName
	claims["role"] = user.Role
	claims["iat"] = time.Now().Unix()
	claims["ext"] = time.Now().Add(30 * time.Minute).Unix()

	cfg := config.Load()

	access, err := accessToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["user_id"] = user.Id
	rftClaims["first_name"] = user.FirstName
	rftClaims["role"] = user.Role
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["ext"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	refresh, err := refreshToken.SignedString([]byte(cfg.REFRESH_SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	t := time.Now().Add(30*time.Minute)
	time := t.Format("2006-01-02 15:04:05")

	return &model.Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresAt:    time,
	}
}

func GenerateAccessToken(user *jwt.MapClaims) *string {

	accessToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = (*user)["user_id"]
	claims["name"] = (*user)["name"]
	claims["nativeLanguage"] = (*user)["nativeLanguage"]
	claims["iat"] = time.Now().Unix()
	claims["ext"] = time.Now().Add(time.Hour).Unix()

	cfg := config.Load()

	access, err := accessToken.SignedString([]byte(cfg.SIGNING_KEY))
	if err != nil {
		log.Fatalf("Access token is not generated %v", err)
	}

	return &access
}

func ExtractClaims(tokenStr string, isRefresh bool) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		if isRefresh {
			return []byte(config.Load().REFRESH_SIGNING_KEY), nil
		}
		return []byte(config.Load().SIGNING_KEY), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
