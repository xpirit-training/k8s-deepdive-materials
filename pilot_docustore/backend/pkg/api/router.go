package api

import (
	"github.com/gorilla/mux"
	"github.com/pathogende/docustore/pkg/database"
	"github.com/sirupsen/logrus"
)

// Router sets up API routes and handlers
func Router(db *database.DB, logger *logrus.Logger) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/podinfo/", podinfo(logger)).Methods("GET")
	r.HandleFunc("/api/v1/document/{id}", getDocumentHandler(db, logger)).Methods("GET")
	r.HandleFunc("/api/v1/document/{id}", updateDocumentHandler(db, logger)).Methods("POST")
	r.HandleFunc("/api/v1/document/{id}", deleteDocumentHandler(db, logger)).Methods("DELETE")
	r.HandleFunc("/api/v1/documents/", listDocumentHandler(db, logger)).Methods("GET")
	r.HandleFunc("/readiness", readiness(logger)).Methods("GET")
	r.HandleFunc("/liveness", liveness(db, logger)).Methods("GET")

	return r
}
