package middleware

import (
	"github.com/gin-gonic/gin"
	"Clean-Grpc/pkg/helper"
)

// User authentcation
func AuthUser(ctx *gin.Context) {
	helper.AuthHelperUser(ctx, "user")
}
