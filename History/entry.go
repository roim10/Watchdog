package history

import (
	"time"

	"github.com/roim10/Watchdog/sources"
)

type Entry struct {
	Result sources.Result
	Time   time.Time
}
