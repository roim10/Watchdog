package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	read "github.com/roim10/Watchdog/Read"
	sourcestore "github.com/roim10/Watchdog/sourceStore"
)

func DeleteSourceHandler(store *sourcestore.SourceStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		err := store.Delete(name)
		if err != nil {
			http.Error(w, "source does not found", http.StatusNotFound)
			return
		}
		err = read.RemoveSource(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "источник успешно удален")
	}
}
func AddSourceHandler(store *sourcestore.SourceStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddSourceRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		if req.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}
		newRead := read.Source{
			Name: req.Name,
			Url:  req.Url,
		}
		err = store.Add(newRead)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		err = read.AppendSource(newRead)
		if err != nil {
			http.Error(w, "Failed to save the data", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Все успешно отработало")
	}
}
