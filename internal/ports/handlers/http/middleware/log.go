package middleware

import (
	"context"
	"net/http"

	"github.com/fakovacic/ports/internal/ports"
	"github.com/google/uuid"
)

func ReqID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, ports.ContextKeyRequestID, uuid.New().String())
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
