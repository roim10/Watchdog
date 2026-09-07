package sources

type Result struct {
	Code int    `json:"code"`
	Body string `json:"body"`
	Err  error  `json:"-"`
}
