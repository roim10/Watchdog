package main

import (
	"fmt"
	"log"
	"sync"

	read "github.com/roim10/Watchdog/Read"
	"github.com/roim10/Watchdog/registry"
	"github.com/roim10/Watchdog/sources"
)

func main() {
	reg := registry.New()
	lines, err := read.Read()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for _, v := range lines {
		wg.Go(func() {
			r := sources.FetchAndPrint(v.Url)
			reg.Set(v.Name, r)
		})
	}
	wg.Wait()
	all := reg.GetAll()
	for name, r := range all {
		if r.Err != nil {
			fmt.Println(name, "ОШИБКА:", r.Err)
		} else {
			fmt.Println(name, "OK", r.Code, r.Body)
		}
	}
}
