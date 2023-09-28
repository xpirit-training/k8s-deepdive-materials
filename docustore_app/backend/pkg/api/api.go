package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pathogende/docustore/pkg/database"
)

func createDocumentHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var document database.Document
		err := json.NewDecoder(r.Body).Decode(&document)
		if err != nil {
			http.Error(w, "Error decoding JSON request: "+err.Error(), http.StatusBadRequest)
			return
		}
		id, err := db.CreateDocument(document)
		if err != nil {
			http.Error(w, "Error creating document: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(id))
	}
}

func getDocumentHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		document, err := db.GetDocument(id)
		if err != nil {
			if err == database.ErrDocumentNotFound {
				http.Error(w, "Document not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error getting document: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(document)
	}
}

func updateDocumentHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var document database.Document
		err := json.NewDecoder(r.Body).Decode(&document)
		if err != nil {
			http.Error(w, "Error decoding JSON request: "+err.Error(), http.StatusBadRequest)
			return
		}
		err = db.UpdateDocument(id, document)
		if err != nil {
			if err == database.ErrDocumentNotFound {
				http.Error(w, "Document not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error updating document: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func deleteDocumentHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := db.DeleteDocument(id)
		if err == database.ErrDocumentNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
