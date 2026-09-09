package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roim10/Watchdog/registry"
)

func NameHandle(reg *registry.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		result, ok := reg.Get(name)
		if !ok {
			http.Error(w, "source with this name not found", http.StatusNotFound)
			return
		}
		status := toSourceStatus(name, result)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status)
	}
}
