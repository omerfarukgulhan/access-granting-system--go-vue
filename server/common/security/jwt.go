package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var (
	ErrTokenExpired        = errors.New("token expired")
	ErrTokenInvalid        = errors.New("invalid token")
	ErrTokenSignatureError = errors.New("invalid token signature")
	ErrTokenClaimsInvalid  = errors.New("invalid token claims")
)

var jwtSecretKey = []byte(loadJwtSecretKey())

type Claims struct {
	UserID int64    `json:"user_id"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

func loadJwtSecretKey() string {
	if err := godotenv.Load(); err != nil {
		return fmt.Sprintf("failed to load environment variables: %v", err)
	}

	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		fmt.Println("Warning: JWT_SECRET_KEY is not set, using default key for local development")
	}

	return key
}

func GenerateToken(userID int64, email string, roles []string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "access-granting",
			Audience:  jwt.ClaimStrings{"access-granting-users"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (int64, string, []string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenSignatureError
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, "", nil, ErrTokenExpired
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return 0, "", nil, ErrTokenSignatureError
		}
		return 0, "", nil, fmt.Errorf("failed to validate token: %w", err)
	}

	if !token.Valid {
		return 0, "", nil, ErrTokenInvalid
	}

	if err := validateClaims(claims); err != nil {
		return 0, "", nil, err
	}

	return claims.UserID, claims.Email, claims.Roles, nil
}

func validateClaims(claims *Claims) error {
	if claims.UserID == 0 {
		return ErrTokenClaimsInvalid
	}
	if claims.Email == "" {
		return ErrTokenClaimsInvalid
	}
	return nil
}
