package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/virtual-terminal", app.VirtualTerminal)
	r.Post("/payment-succeeded", app.PaymentSucceed)

	return r
}
