package read

import (
	"fmt"
	"os"
)

func AppendSource(s Source) error {
	file, err := os.OpenFile("api.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("file doesnt exist. %v", err)
	}
	defer file.Close()
	str := fmt.Sprintf("%v %v\n", s.Name, s.Url)
	_, err = file.WriteString(str)
	if err != nil {
		return fmt.Errorf("Write error: %v", err)
	}
	return nil
}

func RemoveSource(name string) error {
	sources, err := Read()
	if err != nil {
		return fmt.Errorf("ошибка чтения источников: %w", err)
	}
	var result []Source
	exist := false
	for _, src := range sources {
		if src.Name == name {
			exist = true
			continue
		}
		result = append(result, src)
	}
	if !exist {
		return fmt.Errorf("source %q does not exist", name)
	}
	file, err := os.OpenFile("api.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	for _, src := range result {
		str := fmt.Sprintf("%s %s\n", src.Name, src.Url)
		_, err := file.WriteString(str)
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
	}
	return nil

}
