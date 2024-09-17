package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/Stern-Ritter/gophertask/internal/model"
)

func GetPasswordHash(password string) (string, error) {
	if len(strings.TrimSpace(password)) == 0 {
		return "", fmt.Errorf("empty password")
	}
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	byteHash := []byte(hash)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	return err == nil
}

func NewToken(user model.User, secret string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["login"] = user.Login
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid jwt token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return nil, fmt.Errorf("token is expired")
			}
		} else {
			return nil, fmt.Errorf("exp claim is missing or invalid")
		}
	} else {
		return nil, fmt.Errorf("invalid jwt token")
	}

	return token, nil
}
