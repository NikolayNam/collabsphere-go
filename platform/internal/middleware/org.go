package middleware

import (
	"context"
	"net/http"
)

// Читает X-Organization-ID и кладёт в context.
// Тебе это нужно под user_organizations_roles модель.
func OrgContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		orgID := r.Header.Get("X-Organization-ID")
		if orgID != "" {
			ctx := context.WithValue(r.Context(), OrganizationIDKey, orgID)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
