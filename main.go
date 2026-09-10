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
	"github.com/roim10/Watchdog/survey"
)

func main() {
	reg := registry.New()
	hist := history.NewHistory(5)
	lines, err := read.Read()
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	go func() {
		survey.Poll(lines, reg, hist)
		for {
			<-ticker.C
			survey.Poll(lines, reg, hist)
		}
	}()
	router := mux.NewRouter()
	router.HandleFunc("/status", httpapi.StatusHandler(reg))
	router.HandleFunc("/status/{name}", httpapi.NameHandle(reg))
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
