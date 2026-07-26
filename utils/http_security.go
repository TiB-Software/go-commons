package utils

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/TB-Systems/go-commons/errors"
	"github.com/gin-gonic/gin"
)

const (
	DefaultCSRFHeaderName = "X-CSRF-Token"
)

type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

type SecurityHeadersConfig struct {
	Production bool
}

type HTTPSConfig struct {
	Production     bool
	TrustedProxies []string
}

type CSRFConfig struct {
	SessionCookieName string
	HeaderName        string
	Validate          func(ctx context.Context, sessionToken string, csrfToken string) errors.ApiError
}

func CORS(cfg CORSConfig) gin.HandlerFunc {
	allowedOrigins := allowedOriginSet(cfg.AllowedOrigins)
	allowedMethods := headerListOrDefault(cfg.AllowedMethods, []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	allowedHeaders := headerListOrDefault(cfg.AllowedHeaders, []string{"Content-Type", DefaultCSRFHeaderName})

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if _, allowed := allowedOrigins[origin]; allowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", allowedMethods)
			c.Header("Access-Control-Allow-Headers", allowedHeaders)
			c.Header("Vary", "Origin")
			if cfg.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SecurityHeaders(cfg SecurityHeadersConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "no-referrer")

		if cfg.Production && IsSecureRequest(c, nil) {
			c.Header("Strict-Transport-Security", "max-age=31536000")
		}

		c.Next()
	}
}

func RequireHTTPS(cfg HTTPSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !cfg.Production || IsSecureRequest(c, cfg.TrustedProxies) {
			c.Next()
			return
		}

		SendErrorResponse(c, errors.NewApiError(
			http.StatusUpgradeRequired,
			errors.BadRequestError("HTTPS_REQUIRED"),
		))
	}
}

func CSRFRequired(cfg CSRFConfig) gin.HandlerFunc {
	headerName := strings.TrimSpace(cfg.HeaderName)
	if headerName == "" {
		headerName = DefaultCSRFHeaderName
	}

	return func(c *gin.Context) {
		if IsCSRFSafeMethod(c.Request.Method) {
			c.Next()
			return
		}

		sessionToken, err := c.Cookie(cfg.SessionCookieName)
		if err != nil {
			SendErrorResponse(c, errors.NewApiError(
				http.StatusUnauthorized,
				errors.BadRequestError("INVALID_SESSION"),
			))
			return
		}

		if cfg.Validate == nil {
			SendErrorResponse(c, invalidCSRFError())
			return
		}

		csrfToken := c.GetHeader(headerName)
		if apiErr := cfg.Validate(c.Request.Context(), sessionToken, csrfToken); apiErr != nil {
			SendErrorResponse(c, apiErr)
			return
		}

		c.Next()
	}
}

func IsCSRFSafeMethod(method string) bool {
	return method == http.MethodGet || method == http.MethodHead || method == http.MethodOptions
}

func IsSecureRequest(c *gin.Context, trustedProxies []string) bool {
	if c.Request.TLS != nil {
		return true
	}

	if !strings.EqualFold(c.GetHeader("X-Forwarded-Proto"), "https") {
		return false
	}

	if len(trustedProxies) == 0 {
		return true
	}

	return IsTrustedProxy(c.Request.RemoteAddr, trustedProxies)
}

func IsTrustedProxy(remoteAddr string, trustedProxies []string) bool {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		host = remoteAddr
	}

	remoteIP := net.ParseIP(host)
	if remoteIP == nil {
		return false
	}

	for _, trustedProxy := range trustedProxies {
		trustedProxy = strings.TrimSpace(trustedProxy)
		if trustedProxy == "" {
			continue
		}

		if strings.Contains(trustedProxy, "/") {
			_, network, err := net.ParseCIDR(trustedProxy)
			if err == nil && network.Contains(remoteIP) {
				return true
			}
			continue
		}

		trustedIP := net.ParseIP(trustedProxy)
		if trustedIP != nil && trustedIP.Equal(remoteIP) {
			return true
		}
	}

	return false
}

func allowedOriginSet(origins []string) map[string]struct{} {
	result := make(map[string]struct{}, len(origins))
	for _, origin := range origins {
		origin = strings.TrimSpace(origin)
		if origin != "" && origin != "*" {
			result[origin] = struct{}{}
		}
	}

	return result
}

func headerListOrDefault(values []string, fallback []string) string {
	filtered := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			filtered = append(filtered, value)
		}
	}

	if len(filtered) == 0 {
		filtered = fallback
	}

	return strings.Join(filtered, ",")
}

func invalidCSRFError() errors.ApiError {
	return errors.NewApiError(
		http.StatusForbidden,
		errors.BadRequestError("INVALID_CSRF_TOKEN"),
	)
}
