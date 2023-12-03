package code

import (
	"bufio"
	"fmt"
	"os"
)

func Day3() {
	fmt.Println("Starting Day1...")

	file, err := readFile("inputs/day3.txt")
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range file {
		fmt.Println(line)
	}
}

func readFile(name string) ([]string, error) {
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

func findNumbers(data []string) int {
}
