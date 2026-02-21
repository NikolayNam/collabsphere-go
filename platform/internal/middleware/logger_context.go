package middleware

import (
	"net/http"

	"log/slog"

	chimw "github.com/go-chi/chi/v5/middleware"

	appLogger "github.com/NikolayNam/collabsphere-go/internal/logger"
)

func LoggerContext(base *slog.Logger, orgIDHeader string) func(http.Handler) http.Handler {
	if base == nil {
		panic("base logger is nil")
	}
	if orgIDHeader == "" {
		orgIDHeader = "X-Organization-ID"
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := chimw.GetReqID(r.Context())
			orgID := r.Header.Get(orgIDHeader)

			l := base.With(
				"request_id", reqID,
				"org_id", orgID,
				"method", r.Method,
				"path", r.URL.Path,
			)

			ctx := appLogger.With(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
