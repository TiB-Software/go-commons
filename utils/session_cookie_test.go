package utils

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestSetSessionToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("sets http only lax session cookie with positive max age", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/login", nil)

		SetSessionToken(ctx, "session_cookie", "session-token", time.Now().Add(time.Hour))

		cookie := findCookie(t, w, "session_cookie")
		if cookie.Value != "session-token" {
			t.Errorf("Expected cookie value %q, got %q", "session-token", cookie.Value)
		}
		if cookie.Path != "/" {
			t.Errorf("Expected path %q, got %q", "/", cookie.Path)
		}
		if !cookie.HttpOnly {
			t.Error("Expected cookie to be HttpOnly")
		}
		if cookie.Secure {
			t.Error("Expected cookie to be insecure for an HTTP request")
		}
		if cookie.MaxAge <= 0 || cookie.MaxAge > int(time.Hour.Seconds()) {
			t.Errorf("Expected max age between 1 and 3600, got %d", cookie.MaxAge)
		}

		assertSetCookieHeaderContains(t, w, "SameSite=Lax")
	})

	t.Run("sets expired max age to zero", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/login", nil)

		SetSessionToken(ctx, "session_cookie", "session-token", time.Now().Add(-time.Minute))

		cookie := findCookie(t, w, "session_cookie")
		if cookie.MaxAge != 0 {
			t.Errorf("Expected max age 0, got %d", cookie.MaxAge)
		}
	})

	t.Run("sets secure cookie for TLS request", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/login", nil)
		ctx.Request.TLS = &tls.ConnectionState{}

		SetSessionToken(ctx, "session_cookie", "session-token", time.Now().Add(time.Hour))

		cookie := findCookie(t, w, "session_cookie")
		if !cookie.Secure {
			t.Error("Expected cookie to be Secure for a TLS request")
		}
	})

	t.Run("sets secure cookie for forwarded https request", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/login", nil)
		ctx.Request.Header.Set("X-Forwarded-Proto", "HTTPS")

		SetSessionToken(ctx, "session_cookie", "session-token", time.Now().Add(time.Hour))

		cookie := findCookie(t, w, "session_cookie")
		if !cookie.Secure {
			t.Error("Expected cookie to be Secure for a forwarded HTTPS request")
		}
	})
}

func TestClearSessionToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("clears http only lax session cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/logout", nil)

		ClearSessionToken(ctx, "session_cookie")

		cookie := findCookie(t, w, "session_cookie")
		if cookie.Value != "" {
			t.Errorf("Expected cookie value to be blank, got %q", cookie.Value)
		}
		if cookie.MaxAge != -1 {
			t.Errorf("Expected max age -1, got %d", cookie.MaxAge)
		}
		if cookie.Path != "/" {
			t.Errorf("Expected path %q, got %q", "/", cookie.Path)
		}
		if !cookie.HttpOnly {
			t.Error("Expected cookie to be HttpOnly")
		}
		if cookie.Secure {
			t.Error("Expected cookie to be insecure for an HTTP request")
		}

		assertSetCookieHeaderContains(t, w, "SameSite=Lax")
	})

	t.Run("clears secure cookie for forwarded https request", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/logout", nil)
		ctx.Request.Header.Set("X-Forwarded-Proto", "https")

		ClearSessionToken(ctx, "session_cookie")

		cookie := findCookie(t, w, "session_cookie")
		if !cookie.Secure {
			t.Error("Expected cookie to be Secure for a forwarded HTTPS request")
		}
	})
}

func TestSetCSRFToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("sets readable lax csrf cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/login", nil)

		SetCSRFToken(ctx, "csrf_cookie", "csrf-token", time.Now().Add(time.Hour))

		cookie := findCookie(t, w, "csrf_cookie")
		if cookie.Value != "csrf-token" {
			t.Errorf("Expected cookie value %q, got %q", "csrf-token", cookie.Value)
		}
		if cookie.HttpOnly {
			t.Error("Expected CSRF cookie not to be HttpOnly")
		}
		if cookie.MaxAge <= 0 || cookie.MaxAge > int(time.Hour.Seconds()) {
			t.Errorf("Expected max age between 1 and 3600, got %d", cookie.MaxAge)
		}
		assertSetCookieHeaderContains(t, w, "SameSite=Lax")
	})

	t.Run("sets expired max age to zero", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, "/login", nil)

		SetCSRFToken(ctx, "csrf_cookie", "csrf-token", time.Now().Add(-time.Minute))

		cookie := findCookie(t, w, "csrf_cookie")
		if cookie.MaxAge != 0 {
			t.Errorf("Expected max age 0, got %d", cookie.MaxAge)
		}
	})
}

func TestClearCSRFToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest(http.MethodPost, "/logout", nil)

	ClearCSRFToken(ctx, "csrf_cookie")

	cookie := findCookie(t, w, "csrf_cookie")
	if cookie.Value != "" {
		t.Errorf("Expected cookie value to be blank, got %q", cookie.Value)
	}
	if cookie.MaxAge != -1 {
		t.Errorf("Expected max age -1, got %d", cookie.MaxAge)
	}
	if cookie.HttpOnly {
		t.Error("Expected CSRF cookie not to be HttpOnly")
	}
	assertSetCookieHeaderContains(t, w, "SameSite=Lax")
}

func TestIsSecureRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name              string
		tls               bool
		forwardedProtocol string
		expected          bool
	}{
		{
			name:     "tls request",
			tls:      true,
			expected: true,
		},
		{
			name:              "forwarded https request",
			forwardedProtocol: "https",
			expected:          true,
		},
		{
			name:              "forwarded https is case insensitive",
			forwardedProtocol: "HTTPS",
			expected:          true,
		},
		{
			name:              "forwarded http request",
			forwardedProtocol: "http",
			expected:          false,
		},
		{
			name:     "plain request",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.tls {
				ctx.Request.TLS = &tls.ConnectionState{}
			}
			if tt.forwardedProtocol != "" {
				ctx.Request.Header.Set("X-Forwarded-Proto", tt.forwardedProtocol)
			}

			got := isSecureRequest(ctx)
			if got != tt.expected {
				t.Errorf("Expected secure request to be %v, got %v", tt.expected, got)
			}
		})
	}
}

func findCookie(t *testing.T, w *httptest.ResponseRecorder, name string) *http.Cookie {
	t.Helper()

	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}

	t.Fatalf("Expected cookie %q to be set", name)
	return nil
}

func assertSetCookieHeaderContains(t *testing.T, w *httptest.ResponseRecorder, want string) {
	t.Helper()

	for _, header := range w.Result().Header.Values("Set-Cookie") {
		if strings.Contains(header, want) {
			return
		}
	}

	t.Fatalf("Expected Set-Cookie header to contain %q, got %q", want, w.Result().Header.Values("Set-Cookie"))
}
