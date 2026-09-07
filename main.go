package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	read "github.com/roim10/Watchdog/Read"
	"github.com/roim10/Watchdog/registry"
	"github.com/roim10/Watchdog/survey"
)

func main() {
	reg := registry.New()

	lines, err := read.Read()
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	go func() {
		survey.Poll(lines, reg)
		for {
			<-ticker.C
			survey.Poll(lines, reg)
		}
	}()
	router := mux.NewRouter()
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
