package handlers

import (
	"context"
	"duepe/internal/lib/session"
	"duepe/web"
	"log/slog"
	"net/http"
	"strings"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := web.GetPublicFs().ReadFile("static/index.html")
		if err != nil {
			slog.Error("error reading static index file: ", slog.String("err", err.Error()))

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		content = putSessionToIndexContent(r.Context(), content)

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(content)
		if err != nil {
			slog.Error("error writing static index file: ", slog.String("err", err.Error()))
		}
	}
}

func putSessionToIndexContent(ctx context.Context, content []byte) []byte {
	var token string

	ses, ok := session.GetSessionFromContext(ctx)
	if ok {
		token = ses.Token()
	}

	return []byte(strings.Replace(string(content), "{{ .SessionToken }}", token, 1))
}
