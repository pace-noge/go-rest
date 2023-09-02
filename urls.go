package gorest

import "github.com/go-chi/chi/v5"

func registerInternalUrls(r *chi.Mux) {

	r.Get("/health", Health)
}
