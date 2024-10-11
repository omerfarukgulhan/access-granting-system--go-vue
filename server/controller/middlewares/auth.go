package middlewares

import (
	"access-granting/common/security"
	"access-granting/common/util/result"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeaderMissing     = "Authorization header is missing"
	tokenMissing          = "Token is missing"
	tokenExpired          = "Token has expired"
	tokenInvalidOrExpired = "Token is invalid or expired"
	tokenSignatureInvalid = "Token signature is invalid"
	tokenClaimsInvalid    = "Token claims are invalid"
	unauthorizedRole      = "User does not have the required role"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		context.AbortWithStatusJSON(http.StatusUnauthorized, result.NewResult(false, authHeaderMissing))
		return
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, result.NewResult(false, tokenMissing))
		return
	}

	userId, userEmail, userRoles, err := security.ValidateToken(token)
	if err != nil {
		handleTokenError(context, err)
		return
	}

	context.Set("userId", userId)
	context.Set("userEmail", userEmail)
	context.Set("roles", userRoles)

	context.Next()
}

func Authorize(requiredRoles ...string) gin.HandlerFunc {
	return func(context *gin.Context) {
		roles, exists := context.Get("roles")
		if !exists {
			context.AbortWithStatusJSON(http.StatusUnauthorized, result.NewResult(false, tokenClaimsInvalid))
			return
		}

		userRoles, ok := roles.([]string)
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnauthorized, result.NewResult(false, tokenClaimsInvalid))
			return
		}

		if len(userRoles) == 0 {
			context.AbortWithStatusJSON(http.StatusForbidden, result.NewResult(false, "User does not have authorization to perform this action"))
			return
		}

		roleMap := make(map[string]struct{})
		for _, userRole := range userRoles {
			roleMap[userRole] = struct{}{}
		}

		for _, requiredRole := range requiredRoles {
			if _, exists := roleMap[requiredRole]; exists {
				context.Next()
				return
			}
		}

		context.AbortWithStatusJSON(http.StatusForbidden, result.NewResult(false, "User does not have authorization to perform this action"))
	}
}

func handleTokenError(context *gin.Context, err error) {
	var errorMessage string
	switch {
	case errors.Is(err, security.ErrTokenExpired):
		errorMessage = tokenExpired
	case errors.Is(err, security.ErrTokenInvalid):
		errorMessage = tokenInvalidOrExpired
	case errors.Is(err, security.ErrTokenSignatureError):
		errorMessage = tokenSignatureInvalid
	case errors.Is(err, security.ErrTokenClaimsInvalid):
		errorMessage = tokenClaimsInvalid
	default:
		errorMessage = tokenInvalidOrExpired
	}

	context.AbortWithStatusJSON(http.StatusUnauthorized, result.NewResult(false, errorMessage))
}
