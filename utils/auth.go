package utils

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenMetadata struct {
	Expires int64
}

// GenerateNewAccessToken returns JWT token
func GenerateNewAccessToken() (string, error) {
	// get the JWT secret key from .env file
	secret := GetValue("JWT_SECRET_KEY")

	// get the JWT token expire time from .env file
	minutesCount, _ := strconv.Atoi(GetValue("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// create a JWT claim object
	claims := jwt.MapClaims{}

	// add expiration time for the token
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// create a new JWT token with the JWT claim object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// convert the token in a string format
	t, err := token.SignedString([]byte(secret))

	// if conversion failed, return the error
	if err != nil {
		return "", err
	}

	// return the token
	return t, nil
}

// ExtractTokenMetadata returns token metadata
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	// verify the token
	token, err := verifyToken(c)

	// if verification is failed, return an error
	if err != nil {
		return nil, err
	}

	// get the token claim data
	claims, ok := token.Claims.(jwt.MapClaims)

	//  if token claim data exists and token is valid
	if ok && token.Valid {
		// set the token expiration date
		expires := int64(claims["exp"].(float64))

		// return the token metadata
		return &TokenMetadata{
			Expires: expires,
		}, nil
	}

	// return an error if token is invalid
	return nil, err
}

// CheckToken returns token check result
func CheckToken(c *fiber.Ctx) (bool, error) {
	// get the current time
	now := time.Now().Unix()

	// get the token claim data
	claims, err := ExtractTokenMetadata(c)

	// if claim data is not found or invalid
	// return false
	if err != nil {
		return false, err
	}

	// get the expiration time from the claim data
	expires := claims.Expires

	// if the token is expired
	// return false
	if now > expires {
		return false, err
	}

	// return true, this means the token is valid
	return true, nil
}

// extractToken returns token from the Authorization header
func extractToken(c *fiber.Ctx) string {
	// get the bearer token from the Authorization header
	bearToken := c.Get("Authorization")

	// get the JWT token from the bearer
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		// return the JWT token
		return onlyToken[1]
	}

	// return empty if bearer token is empty
	return ""
}

// verifyToken returns verification result
func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	// get the token from the bearer token
	tokenString := extractToken(c)

	// verify the token with the JWT secret key
	token, err := jwt.Parse(tokenString, jwtKeyFunc)

	// if verification is failed, return an error
	if err != nil {
		return nil, err
	}

	// return the valid token
	return token, nil
}

// jwtKeyFunc returns the JWT secret key
// this function is used to verify the token
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(GetValue("JWT_SECRET_KEY")), nil
}
