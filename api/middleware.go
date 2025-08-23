package api

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/MustafaKheda/simplebank/token"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader(AuthorizationHeaderKey)
		if len(authHeader) == 0 {
			err := errors.New("authorization header is required")
			// If the Authorization header is not provided, return an error
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authHeader)
		log.Default().Println("Authorization header fields:", fields)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			// If the Authorization header format is invalid, return an error
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		// Check if the token type is Bearer
		if strings.ToLower(fields[0]) != AuthorizationTypeBearer {
			err := errors.New("unsupported authorization type")
			// If the token type is not Bearer, return an error
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return

		}
		accessToken := fields[1]
		// Verify the token
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// Store the payload in the context for later use
		c.Set(AuthorizationPayloadKey, payload)
		c.Next()
	}
}
