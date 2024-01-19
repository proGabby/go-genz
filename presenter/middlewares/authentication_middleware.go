package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/proGabby/4genz/data/repo_impl/user_repo_impl.go"
	"github.com/proGabby/4genz/domain/entity"
	"github.com/proGabby/4genz/utils"
)

// AuthMiddleware handles user authentication.
type AuthMiddleware struct {
	userRepo user_repo_impl.UserRepositoryImpl
}

// NewAuthMiddleware creates a new AuthMiddleware instance.
func NewAuthMiddleware(userRepo user_repo_impl.UserRepositoryImpl) *AuthMiddleware {
	return &AuthMiddleware{userRepo: userRepo}
}

// Authenticate is the middleware function that performs user authentication.
func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request header or query parameter
		token := extractToken(r)

		// Verify the token against the user store
		user, err := m.verifyJWTToken(token)
		if err != nil {
			jsonResponse := map[string]interface{}{
				"error": "Unauthorized",
			}

			utils.HandleError(jsonResponse, http.StatusUnauthorized, w)
			return
		}

		// Create a context with the user information
		ctx := context.WithValue(r.Context(), "user", user)

		// Call the next handler with the updated context
		next(w, r.WithContext(ctx))
	}
}

func (m *AuthMiddleware) GenerateJWTToken(user *entity.User) (string, error) {

	secrtKey, err := getJWTSecretKey()

	if err != nil {
		return "", err
	}

	// Define the expiration time for the token (e.g., 1 hour)
	expirationTime := time.Now().Add(12 * time.Hour)

	// Create the JWT claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   fmt.Sprintf("%d", user.Id),
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte(secrtKey) // Replace with a secure secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getJWTSecretKey() (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET_KEY not set")
	}
	return secretKey, nil
}

func (m *AuthMiddleware) verifyJWTToken(tokenString string) (*entity.User, error) {

	secretKey, err := getJWTSecretKey()

	if err != nil {
		return nil, err
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Validate the token
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return nil, err
	}

	user, err := m.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func extractToken(r *http.Request) string {

	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// Token is expected in the format "Bearer <token>"
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) == 2 && strings.ToLower(splitToken[0]) == "bearer" {
			return splitToken[1]
		}
	}
	return r.URL.Query().Get("token")
}
