package utils

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TB-Systems/go-commons/errors"
	"github.com/gin-gonic/gin"
)

func TestCORS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("allows configured origin with credentials", func(t *testing.T) {
		router := corsRouter(CORSConfig{
			AllowedOrigins:   []string{"http://localhost:5173"},
			AllowCredentials: true,
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Origin", "http://localhost:5173")

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
		if got := w.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
			t.Fatalf("Expected allowed origin header, got %q", got)
		}
		if got := w.Header().Get("Access-Control-Allow-Credentials"); got != "true" {
			t.Fatalf("Expected credentials true, got %q", got)
		}
		if got := w.Header().Get("Vary"); got != "Origin" {
			t.Fatalf("Expected Vary Origin, got %q", got)
		}
	})

	t.Run("does not allow unconfigured origin", func(t *testing.T) {
		router := corsRouter(CORSConfig{AllowedOrigins: []string{"http://localhost:5173"}})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Origin", "https://evil.example")

		router.ServeHTTP(w, req)

		if got := w.Header().Get("Access-Control-Allow-Origin"); got != "" {
			t.Fatalf("Expected no allowed origin header, got %q", got)
		}
	})

	t.Run("handles preflight with custom methods and headers", func(t *testing.T) {
		router := corsRouter(CORSConfig{
			AllowedOrigins: []string{"http://localhost:5173"},
			AllowedMethods: []string{"POST"},
			AllowedHeaders: []string{"Content-Type", "X-Test"},
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodOptions, "/", nil)
		req.Header.Set("Origin", "http://localhost:5173")
		req.Header.Set("Access-Control-Request-Method", http.MethodPost)

		router.ServeHTTP(w, req)

		if w.Code != http.StatusNoContent {
			t.Fatalf("Expected status %d, got %d", http.StatusNoContent, w.Code)
		}
		if got := w.Header().Get("Access-Control-Allow-Methods"); got != "POST" {
			t.Fatalf("Expected allowed methods POST, got %q", got)
		}
		if got := w.Header().Get("Access-Control-Allow-Headers"); got != "Content-Type,X-Test" {
			t.Fatalf("Expected custom allowed headers, got %q", got)
		}
	})

	t.Run("ignores wildcard origin", func(t *testing.T) {
		router := corsRouter(CORSConfig{AllowedOrigins: []string{"*"}})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Origin", "http://localhost:5173")

		router.ServeHTTP(w, req)

		if got := w.Header().Get("Access-Control-Allow-Origin"); got != "" {
			t.Fatalf("Expected wildcard not to be reflected, got %q", got)
		}
	})
}

func TestSecurityHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("sets common security headers", func(t *testing.T) {
		router := securityHeadersRouter(SecurityHeadersConfig{})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		router.ServeHTTP(w, req)

		if got := w.Header().Get("X-Content-Type-Options"); got != "nosniff" {
			t.Fatalf("Expected nosniff header, got %q", got)
		}
		if got := w.Header().Get("X-Frame-Options"); got != "DENY" {
			t.Fatalf("Expected DENY frame header, got %q", got)
		}
		if got := w.Header().Get("Referrer-Policy"); got != "no-referrer" {
			t.Fatalf("Expected no-referrer policy, got %q", got)
		}
		if got := w.Header().Get("Strict-Transport-Security"); got != "" {
			t.Fatalf("Expected no HSTS outside production, got %q", got)
		}
	})

	t.Run("sets hsts in production for https request", func(t *testing.T) {
		router := securityHeadersRouter(SecurityHeadersConfig{Production: true})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.TLS = &tls.ConnectionState{}

		router.ServeHTTP(w, req)

		if got := w.Header().Get("Strict-Transport-Security"); got != "max-age=31536000" {
			t.Fatalf("Expected HSTS header, got %q", got)
		}
	})
}

func TestRequireHTTPS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("allows http outside production", func(t *testing.T) {
		router := requireHTTPSRouter(HTTPSConfig{})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("rejects http in production", func(t *testing.T) {
		router := requireHTTPSRouter(HTTPSConfig{Production: true})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		router.ServeHTTP(w, req)

		if w.Code != http.StatusUpgradeRequired {
			t.Fatalf("Expected status %d, got %d", http.StatusUpgradeRequired, w.Code)
		}
	})

	t.Run("allows trusted forwarded https in production", func(t *testing.T) {
		router := requireHTTPSRouter(HTTPSConfig{Production: true, TrustedProxies: []string{"192.0.2.1"}})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = "192.0.2.1:1234"
		req.Header.Set("X-Forwarded-Proto", "https")

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
	})
}

func TestCSRFRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("allows safe method without validation", func(t *testing.T) {
		called := false
		router := csrfRouter(CSRFConfig{
			SessionCookieName: "session",
			Validate: func(context.Context, string, string) errors.ApiError {
				called = true
				return nil
			},
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
		if called {
			t.Fatal("Expected CSRF validation not to be called for safe method")
		}
	})

	t.Run("rejects unsafe method without session cookie", func(t *testing.T) {
		router := csrfRouter(CSRFConfig{SessionCookieName: "session"})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)

		router.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Fatalf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})

	t.Run("rejects unsafe method without validator", func(t *testing.T) {
		router := csrfRouter(CSRFConfig{SessionCookieName: "session"})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(&http.Cookie{Name: "session", Value: "session-token"})

		router.ServeHTTP(w, req)

		if w.Code != http.StatusForbidden {
			t.Fatalf("Expected status %d, got %d", http.StatusForbidden, w.Code)
		}
	})

	t.Run("uses configured header name and validator", func(t *testing.T) {
		var gotSessionToken string
		var gotCSRFToken string
		router := csrfRouter(CSRFConfig{
			SessionCookieName: "session",
			HeaderName:        "X-Test-CSRF",
			Validate: func(_ context.Context, sessionToken string, csrfToken string) errors.ApiError {
				gotSessionToken = sessionToken
				gotCSRFToken = csrfToken
				return nil
			},
		})
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.AddCookie(&http.Cookie{Name: "session", Value: "session-token"})
		req.Header.Set("X-Test-CSRF", "csrf-token")

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
		if gotSessionToken != "session-token" {
			t.Fatalf("Expected session token to be passed, got %q", gotSessionToken)
		}
		if gotCSRFToken != "csrf-token" {
			t.Fatalf("Expected CSRF token to be passed, got %q", gotCSRFToken)
		}
	})
}

func TestIsSecureRequestWithTrustedProxies(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	ctx.Request.RemoteAddr = "10.1.2.3:1234"
	ctx.Request.Header.Set("X-Forwarded-Proto", "https")

	if !IsSecureRequest(ctx, []string{"10.0.0.0/8"}) {
		t.Fatal("Expected trusted forwarded HTTPS request to be secure")
	}
	if IsSecureRequest(ctx, []string{"127.0.0.1"}) {
		t.Fatal("Expected untrusted forwarded HTTPS request not to be secure")
	}
}

func TestIsTrustedProxy(t *testing.T) {
	tests := []struct {
		name           string
		remoteAddr     string
		trustedProxies []string
		expected       bool
	}{
		{name: "exact trusted proxy", remoteAddr: "127.0.0.1:1234", trustedProxies: []string{"127.0.0.1"}, expected: true},
		{name: "trusted cidr", remoteAddr: "10.1.2.3:1234", trustedProxies: []string{"10.0.0.0/8"}, expected: true},
		{name: "untrusted proxy", remoteAddr: "203.0.113.10:1234", trustedProxies: []string{"127.0.0.1"}, expected: false},
		{name: "invalid remote address", remoteAddr: "not-an-ip", trustedProxies: []string{"127.0.0.1"}, expected: false},
		{name: "ignores invalid trusted proxy values", remoteAddr: "127.0.0.1:1234", trustedProxies: []string{"", "invalid-cidr/33"}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsTrustedProxy(tt.remoteAddr, tt.trustedProxies)
			if got != tt.expected {
				t.Fatalf("Expected trusted proxy to be %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestAllowedOriginSet(t *testing.T) {
	origins := allowedOriginSet([]string{" http://localhost:5173 ", "", "*"})

	if _, ok := origins["http://localhost:5173"]; !ok {
		t.Fatal("Expected trimmed origin to be allowed")
	}
	if _, ok := origins["*"]; ok {
		t.Fatal("Expected wildcard origin to be ignored")
	}
	if _, ok := origins[""]; ok {
		t.Fatal("Expected blank origin to be ignored")
	}
}

func TestHeaderListOrDefault(t *testing.T) {
	got := headerListOrDefault([]string{" A ", "", "B"}, []string{"Fallback"})
	if got != "A,B" {
		t.Fatalf("Expected custom header list A,B, got %q", got)
	}

	got = headerListOrDefault([]string{"", " "}, []string{"Fallback"})
	if got != "Fallback" {
		t.Fatalf("Expected fallback header list, got %q", got)
	}
}

func corsRouter(cfg CORSConfig) *gin.Engine {
	router := gin.New()
	router.Use(CORS(cfg))
	router.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })
	router.OPTIONS("/", func(c *gin.Context) { c.Status(http.StatusOK) })
	return router
}

func securityHeadersRouter(cfg SecurityHeadersConfig) *gin.Engine {
	router := gin.New()
	router.Use(SecurityHeaders(cfg))
	router.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })
	return router
}

func requireHTTPSRouter(cfg HTTPSConfig) *gin.Engine {
	router := gin.New()
	router.Use(RequireHTTPS(cfg))
	router.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })
	return router
}

func csrfRouter(cfg CSRFConfig) *gin.Engine {
	router := gin.New()
	router.Any("/", CSRFRequired(cfg), func(c *gin.Context) { c.Status(http.StatusOK) })
	return router
}
