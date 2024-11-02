package api

import (
	"github.com/ercross/zax/services/utils/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func addRoutes(mux *chi.Mux, logger *log.Logger, accountsService AccountsService, repo *Repository) http.Handler {
	mux.Use(middleware.AllowContentType("application/json"))
	mux.Use(middleware.SetHeader("Content-Type", "application/json"))

	mux.Mount("/admin", adminRoutes(repo, accountsService, logger))
	mux.Get("/health", probeHealth())

	mux.Post("/authenticate-product", authenticateProduct(repo, logger))
	mux.Post("/authenticate-batch", authenticateBatch(repo, logger))
	mux.Post("/report-counterfeit", reportCounterfeit(repo, logger))

	return mux
}

func probeHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func getCounterfeitReportsByLocation(_ *Repository, _ *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func authenticateProduct(_ *Repository, _ *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func authenticateBatch(_ *Repository, _ *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func reportCounterfeit(_ *Repository, _ *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func adminRoutes(repo *Repository, accountsService AccountsService, logger *log.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(adminOnly(accountsService))
	r.Get("/counterfeit-reports", getCounterfeitReportsByLocation(repo, logger))

	return r
}
