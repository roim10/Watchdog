package httpapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
		fmt.Fprintln(w, "источник успешно удален")
	}
}
func AddSourceHandler(store *sourcestore.SourceStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
