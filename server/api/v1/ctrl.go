package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/goliajp/http-api-gin/core"
	"github.com/goliajp/http-api-gin/data"
	"net/http"
	"strings"
)

type (
	ctrl     struct{}
	userCtrl ctrl
	fooCtrl  ctrl
	barCtrl  ctrl
)

var (
	UserCtrl = &userCtrl{}
	FooCtrl  = &fooCtrl{}
	BarCtrl  = &barCtrl{}
)

func AuthUser(ctx *gin.Context) {
	kv := data.GetRedis()

	// get token from header authorization: Bearer {$token}
	token := strings.TrimLeft(ctx.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "unauthorized"})
		return
	}

	// get session by token
	sx, err := core.GetSessionByToken(ctx, kv, token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
		return
	}

	// check expire
	if sx.IsExpired() {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "token expired"})
		return
	}

	// refresh session
	if err := sx.Refresh(ctx, kv); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": "refresh token failed"})
		return
	}

	// set session data to context
	ctx.Set("session", sx.MustJSON())
	ctx.Set("token", token)
	ctx.Set("userId", sx.UserId)
	ctx.Next()
}
