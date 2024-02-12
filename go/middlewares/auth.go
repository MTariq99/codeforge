package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"codeforge/config"
	"codeforge/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	AdminTokenHeaderKey          = "Authorization"
	AdminAuthorizationTypeBearer = "BEARER"
)

func GenerateToken(user *models.LoginUserRes) (string, error) {

	claims := models.UserClaims{
		UserId:    user.UserId,
		UserName:  user.UserName,
		UserEmail: *user.User_email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(config.Cfg.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("error in signed string")
	}

	return tokenStr, nil
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(AdminTokenHeaderKey)
		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "token does not exist in header"})
			return
		}
		tokenStr := strings.Split(token, " ")
		if len(tokenStr) != 2 || strings.ToUpper(tokenStr[0]) != AdminAuthorizationTypeBearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: "invalid authorization format",
			})
			return
		}

		claims, err := ValidateAccessToken(tokenStr[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
			return
		}

		setClaims(c, *claims)
	}
}

func ValidateAccessToken(tokenString string) (*models.UserClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method in auth token")
		}
		return []byte(config.Cfg.JwtSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("error in parsing token claims: %w", err)
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token: token is not valid")
	}

	return claims, nil
}

func setClaims(c *gin.Context, claims models.UserClaims) {
	c.Set("user_id", claims.UserId)
	c.Set("user_email", claims.UserEmail)
}
