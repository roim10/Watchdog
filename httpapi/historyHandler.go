package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	history "github.com/roim10/Watchdog/History"
)

func HistoryHandler(hist *history.History) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		result, ok := hist.GetHistory(name)
		if !ok {
			http.Error(w, "source with this name not found", http.StatusNotFound)
			return
		}
		limitStr := r.URL.Query().Get("limit")
		if limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				http.Error(w, "limit must be an integer", http.StatusBadRequest)
				return
			}
			if limit <= 0 {
				http.Error(w, "limit must be greater than 0", http.StatusBadRequest)
				return
			}
			if limit < len(result) {
				result = result[len(result)-limit:]
			}
		}
		historyEntries := make([]HistoryEntry, 0, len(result))
		for _, entry := range result {
			historyEntries = append(historyEntries, toHistoryEntry(name, entry))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(historyEntries)
	}
}
