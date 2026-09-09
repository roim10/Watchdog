package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/roim10/Watchdog/registry"
)

func StatusHandler(reg *registry.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		all := reg.GetAll()
		result := make([]SourceStatus, 0, len(all))
		for name, res := range all {
			result = append(result, toSourceStatus(name, res))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
