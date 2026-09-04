package sources

import (
	"io"
	"net/http"
	"time"
)

func FetchAndPrint(url string) Result {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Result{Err: err}
	}
	req.Header.Set("User-Agent", "MyCustomApp/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return Result{Err: err}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Err: err}
	}
	result := Result{
		Code: resp.StatusCode,
		Body: string(body),
		Err:  nil}
	return result
}
