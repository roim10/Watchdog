package read

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func Read() ([]string, error) {
	file, err := os.Open("Watchdog/api.txt")
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при сканировании: %w", err)
	}
	if slices.Contains(lines, "") {
		return nil, fmt.Errorf("заполните файл URL-адресами")
	}
	return lines, nil
}
