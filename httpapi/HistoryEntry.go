package httpapi

import (
	"time"

	history "github.com/roim10/Watchdog/History"
)

type HistoryEntry struct {
	SourceStatus
	Time time.Time `json:"time"`
}

func toHistoryEntry(name string, e history.Entry) HistoryEntry {
	status := toSourceStatus(name, e.Result)
	return HistoryEntry{
		SourceStatus: status,
		Time:         e.Time,
	}
}
