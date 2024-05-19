package handlers

import (
	"duepe/web"
	"log/slog"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	content, err := web.GetPublicFs().ReadFile("static/index.html")
	if err != nil {
		slog.Error("error reading static index file: ", slog.String("err", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(content)
	if err != nil {
		slog.Error("error writing static index file: ", slog.String("err", err.Error()))
	}
}
