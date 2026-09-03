package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/roim10/Watchdog/sources"
)

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
	lines = append(lines, urlZero)
	for _, v := range lines {
		code, body, err := sources.FetchAndPrint(v)
		if err != nil {
			log.Println("Не удалось получить данные:", err)
		} else {
			fmt.Println(code, body)
		}
	}
}
