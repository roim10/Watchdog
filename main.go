package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	history "github.com/roim10/Watchdog/History"
	read "github.com/roim10/Watchdog/Read"
	"github.com/roim10/Watchdog/httpapi"
	"github.com/roim10/Watchdog/registry"
	sourcestore "github.com/roim10/Watchdog/sourceStore"
	"github.com/roim10/Watchdog/survey"
)

func main() {
	store := sourcestore.NewSourceStore()
	reg := registry.New()
	hist := history.NewHistory(5)
	lines, err := read.Read()
	if err != nil {
		log.Fatal(err)
	}
	for _, source := range lines {
		if err := store.Add(source); err != nil {
			log.Fatal(err)
		}
	}
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	go func() {
		survey.Poll(store, reg, hist)
		for {
			<-ticker.C
			survey.Poll(store, reg, hist)
		}
	}()
	router := mux.NewRouter()
	router.HandleFunc("/sources", httpapi.AddSourceHandler(store)).Methods("POST")
	router.HandleFunc("/sources/{name}", httpapi.DeleteSourceHandler(store)).Methods("DELETE")
	router.HandleFunc("/history/{name}", httpapi.HistoryHandler(hist))
	router.HandleFunc("/status", httpapi.StatusHandler(reg))
	router.HandleFunc("/status/{name}", httpapi.NameHandle(reg))
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
