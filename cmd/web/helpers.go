package main

import (
	"log/slog"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method    = r.Method
		uri       = r.URL.RequestURI()
		errorCode = http.StatusInternalServerError
	)

	app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
	http.Error(w, http.StatusText(errorCode), errorCode)
}

func (app *application) clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}
