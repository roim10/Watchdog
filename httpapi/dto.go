package httpapi

type SourceStatus struct {
	Name  string `json:"name"`
	Code  int    `json:"code"`
	Body  string `json:"body"`
	Error string `json:"error,omitempty"`
}
