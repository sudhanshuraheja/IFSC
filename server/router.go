package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudhanshuraheja/ifsc/datastore"
	"github.com/sudhanshuraheja/ifsc/logger"
)

type cache struct {
	inline map[string]string
}

var cached cache

// Router : route requests to handlers
func Router() http.Handler {
	cached.inline = make(map[string]string)

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
	vars := mux.Vars(r)
	query := vars["query"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if cachedData, ok := cached.inline[query]; ok {
		w.Write([]byte(cachedData))
	} else {
		search := datastore.Search(query)
		data, err := json.Marshal(search)
		if err != nil {
			logger.Error(err)
		}

		cached.inline[query] = string(data)
		w.Write([]byte(data))
	}
}
