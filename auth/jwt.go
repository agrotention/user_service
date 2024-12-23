package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

// Custom claims struct
type claims struct {
	Exp *jwt.NumericDate `json:"exp"`
	Iat *jwt.NumericDate `json:"iat"`
	Nbf *jwt.NumericDate `json:"nbf"`
	Iss string           `json:"iss"`
	Sub string           `json:"sub"`
	Aud jwt.ClaimStrings `json:"aud"`
}

// Create a new Claims object
func NewClaims(sub string) *claims {
	now := time.Now()
	return &claims{
		Exp: jwt.NewNumericDate(now.AddDate(0, 0, 7)), // Add 7 days
		Iat: jwt.NewNumericDate(now),
		Nbf: jwt.NewNumericDate(now),
		Iss: "api.agrotention.io",
		Sub: sub,
		Aud: []string{"api.agrotention.io"},
	}
}

// Create a JWT token
func CreateToken(claims *claims) (string, error) {
	secret := os.Getenv("APP_SECRET")
	if secret == "" {
		return "", errors.New("APP_SECRET is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret)) // Convert secret to []byte
}

// Parse and validate a JWT token
func ParseToken(tokenString string) (*claims, error) {
	var claims claims

	secret := os.Getenv("APP_SECRET")
	if secret == "" {
		return nil, errors.New("APP_SECRET is not set")
	}

	// Parse token
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil // Return secret as []byte
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}

// Accessors for Claims struct
func (c *claims) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.Exp, nil
}

func (c *claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.Iat, nil
}

func (c *claims) GetNotBefore() (*jwt.NumericDate, error) {
	return c.Nbf, nil
}

func (c *claims) GetIssuer() (string, error) {
	return c.Iss, nil
}

func (c *claims) GetSubject() (string, error) {
	return c.Sub, nil
}

func (c *claims) GetAudience() (jwt.ClaimStrings, error) {
	return c.Aud, nil
}
