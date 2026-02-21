package middleware

import (
	_ "context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"golang.org/x/time/rate"
)

type ctxKey string

const OrganizationIDKey ctxKey = "organization_id"

var limiter = rate.NewLimiter(5, 10) // 5 rps, burst 10

func Stack(next http.Handler) http.Handler {
	return chi.Chain(
		chimw.RequestID,
		chimw.RealIP,
		chimw.Recoverer,
		chimw.Timeout(30*time.Second),
		SecurityHeaders,
		OrgContext,
		RateLimit,
		chimw.RequestSize(1<<20), // 1MB
	).Handler(next)
}
