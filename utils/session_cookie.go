package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SetSessionToken(ctx *gin.Context, cookieName string, token string, expiresAt time.Time) {
	maxAge := int(time.Until(expiresAt).Seconds())
	if maxAge < 0 {
		maxAge = 0
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		cookieName,
		token,
		maxAge,
		"/",
		"",
		isSecureRequest(ctx),
		true,
	)
}

func ClearSessionToken(ctx *gin.Context, cookieName string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		cookieName,
		"",
		-1,
		"/",
		"",
		isSecureRequest(ctx),
		true,
	)
}

func SetCSRFToken(ctx *gin.Context, cookieName string, token string, expiresAt time.Time) {
	maxAge := int(time.Until(expiresAt).Seconds())
	if maxAge < 0 {
		maxAge = 0
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		cookieName,
		token,
		maxAge,
		"/",
		"",
		isSecureRequest(ctx),
		false,
	)
}

func ClearCSRFToken(ctx *gin.Context, cookieName string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(
		cookieName,
		"",
		-1,
		"/",
		"",
		isSecureRequest(ctx),
		false,
	)
}

func isSecureRequest(ctx *gin.Context) bool {
	return ctx.Request.TLS != nil || strings.EqualFold(ctx.GetHeader("X-Forwarded-Proto"), "https")
}
