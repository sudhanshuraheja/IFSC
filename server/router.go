package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/ifsc/datastore"
	"github.com/sudhanshuraheja/ifsc/logger"
)

// Router : route requests to handlers
func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/search/{query}", searchHandler).Methods("GET")
	return router
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\":\"pong\"}"))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	query := vars["query"]

	search := datastore.Search(query)
	data, err := json.Marshal(search)
	if err != nil {
		logger.Error(err)
	}
	w.Write([]byte(data))
}
