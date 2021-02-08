package restapi

import (
	"github.com/dgrijalva/jwt-go"
	jwtRequest "github.com/dgrijalva/jwt-go/request"
	"group-management-api/domain/model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// This will be saved inside our token.
type TokenClaims struct {
	UserID         model.UserID       `json:"user_id"`
	StandardClaims jwt.StandardClaims `json:"standard_claims"`
}

func (t TokenClaims) Valid() error {
	return t.StandardClaims.Valid()
}

// Extract the string from Authorization header.
var AuthenticationHeaderExtractorFilter = jwtRequest.PostExtractionFilter{
	Extractor: jwtRequest.HeaderExtractor{"Authorization"},
	Filter: func(s string) (string, error) {
		tokenString := strings.TrimPrefix(s, "Bearer ")
		return tokenString, nil
	}}

func (s *Server) ParseToken(r *http.Request, tokenClaims *TokenClaims) (*jwt.Token, error) {
	token, err := jwtRequest.ParseFromRequest(r, &AuthenticationHeaderExtractorFilter, func(token *jwt.Token) (interface{}, error) {
		key := []byte(s.JwtSecret)
		return key, nil
	}, jwtRequest.WithClaims(tokenClaims))
	return token, err
}

func  (s *Server)  GenerateToken(userID model.UserID) (*string, error) {
	// The tokens will expire in one day. Unix function converts the
	// date to the seconds passed so int64.
	expiresAt := time.Now().Add(time.Hour * 10)

	// Populate the claims.
	claims := TokenClaims{
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.FormatInt(int64(userID), 10),
			Issuer:    "GroupManagementApp",
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the JWT_KEY environment variable.
	key := []byte(s.JwtSecret)
	signedString, err := token.SignedString(key)

	if err != nil {
		return nil, err
	}

	return &signedString, nil
}