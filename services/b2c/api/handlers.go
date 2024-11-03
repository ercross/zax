package api

import "net/http"

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
