package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/roim10/Watchdog/sources"
)

type Result struct {
	Code int
	Body string
	Err  error
}

func main() {
	var wg sync.WaitGroup
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
	lines = append(lines, urlZero)
	result := make(chan Result, len(lines))
	for _, v := range lines {
		wg.Go(func() {
			code, body, err := sources.FetchAndPrint(v)

			if err != nil {
				log.Println("Не удалось получить данные:", err)
			}
			result <- Result{
				Code: code,
				Body: body,
				Err:  err,
			}
		})
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	for r := range result {
		if r.Err != nil {
			fmt.Println("ОШИБКА:", r.Err)
		} else {
			fmt.Println("OK", r.Code, r.Body)
		}
	}
}
