package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MustafaKheda/simplebank/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthorization(t *testing.T, request *http.Request, tokenMaker token.Maker, authorizationType string, username string, duration time.Duration) {
	// This function adds an authorization header to the request.
	// It generates a token for the given username and sets it in the request header.
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(AuthorizationHeaderKey, authorizationHeader)
}
func TestAuthMiddleWare(t *testing.T) {
	// This is a placeholder for middleware tests.
	// You can implement tests for your middleware functions here.
	// For example, you can test the authMiddleware function to ensure it correctly
	// validates tokens and handles unauthorized requests.
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "ValidToken",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				// This function sets up a valid token for the request.
				addAuthorization(t, request, tokenMaker, AuthorizationTypeBearer, "user1", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// This function checks the response for a valid token.
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "UnsupportedAuthrization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				// This function sets up a valid token for the request.
				addAuthorization(t, request, tokenMaker, "unsupported", "user1", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// This function checks the response for a valid token.
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "noAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				// This function sets up a valid token for the request.
				// addAuthorization(t, request, tokenMaker, AuthorizationTypeBearer, "user1", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// This function checks the response for a valid token.
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "invalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				// This function sets up a valid token for the request.
				addAuthorization(t, request, tokenMaker, "", "user1", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// This function checks the response for a valid token.
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				// This function sets up a valid token for the request.
				addAuthorization(t, request, tokenMaker, AuthorizationTypeBearer, "user1", -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// This function checks the response for a valid token.
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)

			authPath := "/auth"
			server.router.GET(authPath, authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					// This is a placeholder for the actual handler logic.
					// You can replace it with your actual handler function.
					ctx.JSON(http.StatusOK, gin.H{})

				},
			)
			// Create a new HTTP request
			recorder := httptest.NewRecorder()
			// This is a placeholder for the actual request setup.
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})

	}
}
