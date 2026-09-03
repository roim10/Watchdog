package sources

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func FetchAndPrint(url string) (int, string, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return 0, "", errors.New("Ошибка создания запроса")
	}
	req.Header.Set("User-Agent", "MyCustomApp/1.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка выполнения запроса: %v", err)
		return 0, "", errors.New("Ошибка выполнения запроса")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка чтения ответа: %v", err)
		return 0, "", errors.New("Ошибка чтения ответа")
	}
	return resp.StatusCode, string(body), nil
}
