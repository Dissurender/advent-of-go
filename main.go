package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Dissurender/advent-of-code/code"
)

type DayFunction func()

func main() {
	// daysFunctions must be manually set
	dayFunctions := map[int]DayFunction{
		1: code.Day1,
		2: code.Day2,
		3: code.Day3,
	}

	dayFuncs, err := countFiles("code")
	if err != nil {
		fmt.Printf("error counting days: %v\n", err)
	}

	funcCount := len(dayFunctions)
	dayCount := dayFuncs
	if dayCount != funcCount {
		fmt.Println("count mismatch")
	}

	fmt.Printf("Select a day 1-%d:\n", dayCount)

	reader := bufio.NewReader(os.Stdin)
	day := -1

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		day, err = strconv.Atoi(input)
		if err != nil {
			fmt.Printf("invalid input. please enter a number 1-%d", dayCount)
			continue
		}
		if day < 1 || day > dayCount {
			fmt.Printf("invalid input. please enter a number 1-%d", dayCount)
			continue
		}
		break
	}

	function, has := dayFunctions[day]
	if has {
		function() // run day function as a pseudo first-class action
	}
}

// countFiles counts the number of Go files in a directory `dir`
func countFiles(dir string) (int, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, file := range files {
		testTest := strings.Split(filepath.Base(file.Name()), "_")

		if filepath.Ext(file.Name()) == ".go" && len(testTest) == 1 {
			count++
		}
	}

	return count, nil
}
