package main

import (
	"fmt"
	"log"
	"sync"

	read "github.com/roim10/Watchdog/Read"
	"github.com/roim10/Watchdog/sources"
)

func main() {
	lines, err := read.Read()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	lines = append(lines, "http://localhost:1")
	result := make(chan sources.Result, len(lines))
	for _, v := range lines {
		wg.Go(func() {
			r := sources.FetchAndPrint(v)
			result <- r
		})
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	for r := range result {
		if r.Err != nil {
			fmt.Println("ОШИБКА:", r.Err)
		} else {
			fmt.Println("OK", r.Code, r.Body)
		}
	}
}
