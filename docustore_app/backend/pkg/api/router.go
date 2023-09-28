package api

import (
	"github.com/gorilla/mux"
	"github.com/pathogende/docustore/pkg/database"
)

// Router sets up API routes and handlers
func Router(db *database.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/documents", createDocumentHandler(db)).Methods("POST")
	r.HandleFunc("/documents/{id}", getDocumentHandler(db)).Methods("GET")
	r.HandleFunc("/documents/{id}", updateDocumentHandler(db)).Methods("PUT")
	r.HandleFunc("/documents/{id}", deleteDocumentHandler(db)).Methods("DELETE")
	return r
}
