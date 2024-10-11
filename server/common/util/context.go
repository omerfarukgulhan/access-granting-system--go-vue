package util

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(ctx *gin.Context) (int64, error) {
	userId, exists := ctx.Get("userId")
	if !exists {
		return 0, errors.New("user id not found in context")
	}

	userIdInt, ok := userId.(int64)
	if !ok {
		return 0, errors.New("user id type assertion failed")
	}

	return userIdInt, nil
}
