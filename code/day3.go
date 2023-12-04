package code

import (
	"fmt"

	"github.com/Dissurender/advent-of-code/utils"
)

func Day3() {
	fmt.Println("Starting Day1...")

	file, err := utils.FileToSlice("inputs/day3.txt")
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range file {
		fmt.Println(line)
	}
}

func findNumbers(data []string) int {
	return 0
}
