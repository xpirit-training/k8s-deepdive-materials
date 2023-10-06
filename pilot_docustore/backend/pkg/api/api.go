package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pathogende/docustore/pkg/database"
	"github.com/sirupsen/logrus"
)

func respondForPreflight(db *database.DB, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the necessary headers for CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin (or specify a domain)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Optionally check the Origin header in the request against your allowed origins

		// Respond with a 200 status for preflight requests
		w.WriteHeader(http.StatusOK)
	}
}

func getDocumentHandler(db *database.DB, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Got get request from %s with UserAgent: %s ", r.RemoteAddr, r.UserAgent())
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins. For production, you should limit this to specific domains.
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		vars := mux.Vars(r)
		id := vars["id"]
		document, err := db.GetDocument(id)
		if err != nil {
			log.Error(err)
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

func updateDocumentHandler(db *database.DB, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Got update request from %s with UserAgent: %s ", r.RemoteAddr, r.UserAgent())
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins. For production, you should limit this to specific domains.
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		vars := mux.Vars(r)
		id := vars["id"]
		var document database.Document
		err := json.NewDecoder(r.Body).Decode(&document)
		if err != nil {
			log.Error(err)
			http.Error(w, "Error decoding JSON request: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = db.UpdateOrCreateDocument(id, document)
		if err != nil {
			log.Error(err)
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

func deleteDocumentHandler(db *database.DB, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Got delete request from %s with UserAgent: %s ", r.RemoteAddr, r.UserAgent())
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins. For production, you should limit this to specific domains.
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		id := mux.Vars(r)["id"]
		err := db.DeleteDocument(id)
		if err == database.ErrDocumentNotFound {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func listDocumentHandler(db *database.DB, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Got list request from %s with UserAgent: %s ", r.RemoteAddr, r.UserAgent())
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins. For production, you should limit this to specific domains.
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//vars := mux.Vars(r)

		ids, err := db.ListDocumentIDs()
		if err != nil {
			log.Error(err)
			if err == database.ErrDocumentNotFound {
				http.Error(w, "Document not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error listing documents: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ids)
	}
}

func podinfo(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Got podinfo request from %s with UserAgent: %s ", r.RemoteAddr, r.UserAgent())
		type PodInfo struct {
			PodName       string `json:"pod_name,omitempty" bson:"_pod_name,omitempty"`
			NameSpaceName string `json:"namespace_name" bson:"namespace_name"`
		}
		var information PodInfo
		information.PodName, _ = os.LookupEnv("MY_POD_NAME")
		information.NameSpaceName, _ = os.LookupEnv("MY_POD_NAMESPACE")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(information)
	}
}

func readiness(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

func liveness(db *database.DB, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if db == nil {
			log.Error("Database is nil")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Database not initialized"))
			return
		}
		err := db.Ping(ctx)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		} else {
			log.Error(err)
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("failed to reach database"))
		}

	}
}
