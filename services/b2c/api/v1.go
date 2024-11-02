package api

import (
	"github.com/ercross/zax/services/utils/log"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func addRoutes(mux *chi.Mux, logger *log.Logger, repo *Repository) http.Handler {
	mux.Get("/health", probeHealth())
	mux.Get("/counterfeit-reports", getCounterfeitReportsByLocation(repo, logger))

	mux.Post("/authenticate-product", authenticateProduct(repo, logger))
	mux.Post("/authenticate-batch", authenticateBatch(repo, logger))
	mux.Post("/report-counterfeit", reportCounterfeit(repo, logger))

	return mux
}

func probeHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func getCounterfeitReportsByLocation(repo *Repository, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func authenticateProduct(repo *Repository, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func authenticateBatch(repo *Repository, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func reportCounterfeit(repo *Repository, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
