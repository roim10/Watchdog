package httpapi

import "github.com/roim10/Watchdog/sources"

type SourceStatus struct {
	Name  string `json:"name"`
	Code  int    `json:"code"`
	Body  string `json:"body"`
	Error string `json:"error,omitempty"`
}

func toSourceStatus(name string, r sources.Result) SourceStatus {
	status := SourceStatus{
		Name: name,
		Code: r.Code,
		Body: r.Body,
	}
	if r.Err != nil {
		status.Error = r.Err.Error()
	}
	return status
}
