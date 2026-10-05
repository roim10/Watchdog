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
