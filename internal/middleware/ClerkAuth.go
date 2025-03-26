package middleware

import (
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/gin-gonic/gin"
)

// ClerkAuth is a middleware that adapts Clerk's WithHeaderAuthorization for Gin
// The client must provide the session token in the Authorization header, otherwise
// the middleware will return a 401 Unauthorized response.
// https://clerk.com/docs/backend-requests/overview
func ClerkAuth() gin.HandlerFunc {
	// This is Clerk's standard HTTP middleware
	clerkMiddleware := clerkhttp.WithHeaderAuthorization()

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Printf("Authorization header: %s", authHeader)

		// Create a custom ResponseWriter that will be used by Clerk's middleware
		w := &responseWriterWrapper{
			ResponseWriter: c.Writer,
			statusCode:     http.StatusOK,
			written:        false,
		}

		// Create a function that will continue the middleware chain
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Update the context in Gin with the one possibly modified by Clerk
			c.Request = r

			// Log successful authentication
			log.Printf("Authentication successful for request: %s %s", r.Method, r.URL.Path)

			c.Next()
		})

		// Run the Clerk middleware
		handler := clerkMiddleware(next)
		handler.ServeHTTP(w, c.Request)

		// If the Clerk middleware wrote a response, stop the Gin chain
		if w.written {
			c.Abort()
		}
	}
}

// responseWriterWrapper wraps a http.ResponseWriter to detect if a response has been written
type responseWriterWrapper struct {
	gin.ResponseWriter
	statusCode int
	written    bool
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.written = true
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterWrapper) Write(b []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(b)
}

// Helper function to get user claims from Gin context
func GetUserClaims(c *gin.Context) (*clerk.SessionClaims, bool) {
	return clerk.SessionClaimsFromContext(c.Request.Context())
}
