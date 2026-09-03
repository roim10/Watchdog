package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func fetchAndPrint(url string) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return
	}
	req.Header.Set("User-Agent", "MyCustomApp/1.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка выполнения запроса: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка чтения ответа: %v", err)
		return
	}
	fmt.Printf("Статус: %v\n", resp.StatusCode)
	fmt.Printf("Тело ответа:\n%s\n", string(body))
}

func main() {
	file, err := os.Open("api.txt")
	if err != nil {
		log.Printf("Ошибка при открытии файла: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка при сканировании: %v", err)
		return
	}
	var urlOne, urlTwo, urlThree string
	if len(lines) == 3 {
		urlOne = lines[0]
		urlTwo = lines[1]
		urlThree = lines[2]
	}
	if urlOne == "" || urlTwo == "" || urlThree == "" {
		log.Println("Ошибка заполните файл с url адресами")
		return
	}
	urlZero := "http://localhost:1"
	fetchAndPrint(urlOne)
	fetchAndPrint(urlTwo)
	fetchAndPrint(urlThree)
	fetchAndPrint(urlZero)

}
