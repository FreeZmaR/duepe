package middlewares

import (
	"duepe/internal/lib/session"
	"net/http"
)

type Session struct {
	manager *session.Manager
}

func NewSession() *Session {
	return &Session{manager: session.GetManager()}
}

func (s *Session) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Session-Token")
		if token == "" {
			next.ServeHTTP(w, r)

			return
		}

		ses, err := s.manager.GetSession(r.Context(), token)
		if err != nil {
			next.ServeHTTP(w, r)

			return
		}

		ctx := s.manager.SaveSessionToContext(r.Context(), ses)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
