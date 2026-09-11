package handlers

import (
	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)
func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/coins", GetCoinBalance)
		
	})

	}
