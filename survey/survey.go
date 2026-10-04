package survey

import (
	"fmt"
	"sync"

	history "github.com/roim10/Watchdog/History"
	"github.com/roim10/Watchdog/registry"
	sourcestore "github.com/roim10/Watchdog/sourceStore"
	"github.com/roim10/Watchdog/sources"
)

func Poll(store *sourcestore.SourceStore, reg *registry.Registry, h *history.History) {
	var wg sync.WaitGroup
	lines := store.List()
	for _, v := range lines {
		wg.Go(func() {
			r := sources.FetchAndPrint(v.Url)
			reg.Set(v.Name, r)
			h.Add(v.Name, r)
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
