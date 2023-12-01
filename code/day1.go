package code

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day1() {
	fmt.Println("Starting Day1...")

	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		fmt.Printf("error reading input file: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0

	for i := 0; i < 1000; i++ {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			return
		}

		fmt.Printf("line %d is: %s", i+1, line)

		value, err := filter(line) // remove chars and inner digits
		if err != nil {
			fmt.Fprintf(os.Stderr, "error converting to int: %v\n", err)
			return
		}

		fmt.Printf("value is: %d\n", value)
		sum = sum + value
	}

	fmt.Printf("The final calibration is: %d\n", sum)
}

var (
	nonNumericRegex = regexp.MustCompile(`[^1-9]+`)
	numWords        = map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
)

func filter(str string) (int, error) {
	newStr := str

	// TODO: word collisions ie: eightwo
	// maybe hard code each iteration?

	for i := 0; i < len(numWords); i++ {
		newStr = strings.ReplaceAll(newStr, numWords[i+1], strconv.Itoa(i+1))
	}

	newStr = nonNumericRegex.ReplaceAllString(newStr, "")

	fmt.Printf("line after replace: %s\n", newStr)

	first := newStr[0:1]
	last := newStr[len(newStr)-1:]

	combined := first + last

	toInt, err := strconv.Atoi(combined)
	if err != nil {
		return 0, err
	}

	return toInt, nil
}
