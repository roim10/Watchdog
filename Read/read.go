package read

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read() ([]Source, error) {
	file, err := os.Open("api.txt")
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
	var result []Source
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, fmt.Errorf("некорректная строка: %q", line)
		}
		result = append(result, Source{
			Url:  parts[1],
			Name: parts[0],
		})
	}
	return result, nil
}
