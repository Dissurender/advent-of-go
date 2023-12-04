package utils

import (
	"bufio"
	"fmt"
	"os"
)

func FileToSlice(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error reading input file '%s': %v", name, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error scanning file '%s': %v", name, err)
	}

	return data, nil
}
