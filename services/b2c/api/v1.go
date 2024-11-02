package api

import (
	"github.com/ercross/zax/services/utils/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewServer(reqLoggerBase *log.Logger, accountsService AccountsService, repo *Repository) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(log.RequestLogger(reqLoggerBase))
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.StripSlashes)
	mux.Use(middleware.SetHeader("Content-Type", "application/json"))

	addRoutes(mux, accountsService, repo)
	return mux
}

func addRoutes(mux *chi.Mux, accountsService AccountsService, repo *Repository) http.Handler {

	mux.Mount("/admin", adminRoutes(repo, accountsService))
	mux.Mount("/submit-with-file", largeRequestRoutes(repo))
	mux.Mount("/authenticate", productAuthenticationRoutes(repo))
	mux.Get("/health", probeHealth())

	return mux
}

func probeHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func getCounterfeitReportsByLocation(_ *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func authenticateProduct(_ *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func authenticateBatch(_ *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func reportCounterfeit(_ *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func adminRoutes(repo *Repository, accountsService AccountsService) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestSize(500 * 1024))
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(adminOnly(accountsService))

	r.Get("/counterfeit/reports", getCounterfeitReportsByLocation(repo))

	return r
}

func largeRequestRoutes(repo *Repository) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("multipart/form-data"))
	r.Use(middleware.RequestSize(20 << 20)) // 20mb

	r.Post("/counterfeit/report", reportCounterfeit(repo))
	return r
}

func productAuthenticationRoutes(repo *Repository) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestSize(500 * 1024))
	r.Use(middleware.AllowContentType("application/json"))

	r.Post("/product", authenticateProduct(repo))
	r.Post("/batch", authenticateBatch(repo))

	return r
}
