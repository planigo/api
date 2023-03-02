package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TokenPayload defines the payload for the token
type TokenPayload struct {
	Id   string
	Role string
}

func GenerateJWT(payload *TokenPayload) string {
	v, err := time.ParseDuration("48h")
	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(v).Unix(),
		"Id":   payload.Id,
		"Role": payload.Role,
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}

	return token
}

func parseJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func VerifyJWT(token string) (*TokenPayload, error) {
	parsed, err := parseJWT(token)
	if err != nil {
		println("Error: parseJWT ", err.Error(), "\n")
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		println("Error: parsed.Claims.(jwt.MapClaims) ", err.Error(), "\n")
		return nil, err
	}
	id, ok := claims["Id"].(string)
	if !ok {
		return nil, errors.New("something went wrong")
	}
	role, ok := claims["Role"].(string)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{Id: id, Role: role}, nil
}
