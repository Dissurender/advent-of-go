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
		line, err := reader.ReadString('\n') // read line from txt file

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
	numWords        = map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	} // mutation map for word form digits
)

// filter reduces the given string
func filter(str string) (int, error) {
	newStr := str

	// word collisions ie: eightwo twone oneight
	// maybe hard code each iteration?
	// newStr = strings.ReplaceAll(newStr, "eightwo", "82")

	for num, word := range numWords {
		newStr = strings.ReplaceAll(newStr, num, word)
	}

	newStr = nonNumericRegex.ReplaceAllString(newStr, "")

	fmt.Printf("line after replace: %s\n", newStr)

	first := newStr[0:1]           // grab first digit
	last := newStr[len(newStr)-1:] // grab last digit

	combined := first + last

	toInt, err := strconv.Atoi(combined)
	if err != nil {
		return 0, err
	}

	return toInt, nil
}
