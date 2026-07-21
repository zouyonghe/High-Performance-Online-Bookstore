package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
	// ErrInvalidClaims means the token claims were missing or malformed.
	ErrInvalidClaims = errors.New("the token claims are missing or malformed")
)

// Context is the context of the JSON web token.
type Context struct {
	ID       uint64
	Username string
	Role     string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we expect.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token, pinning the HMAC signing method to
	// prevent algorithm confusion attacks.
	token, err := jwt.Parse(tokenString, secretFunc(secret),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return ctx, err
	}

	// Read the token if it is valid.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ctx, jwt.ErrTokenInvalidClaims
	}

	// Safely extract the claims, never panic on malformed tokens.
	id, ok := claims["id"].(float64)
	if !ok {
		return ctx, ErrInvalidClaims
	}
	username, ok := claims["username"].(string)
	if !ok {
		return ctx, ErrInvalidClaims
	}
	role, ok := claims["role"].(string)
	if !ok {
		return ctx, ErrInvalidClaims
	}

	ctx.ID = uint64(id)
	ctx.Username = username
	ctx.Role = role
	return ctx, nil
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parse the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	_, err := fmt.Sscanf(header, "Bearer %s", &t)
	if err != nil {
		return nil, err
	}
	return Parse(t, secret)
}

// Sign signs the context with the specified secret.
func Sign(c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	now := time.Now()
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"role":     c.Role,
		"nbf":      now.Unix(),
		"iat":      now.Unix(),
		"exp":      now.Add(time.Hour * 72).Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
