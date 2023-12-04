package code

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Dissurender/advent-of-code/utils"
)

func Day4() {
	file, err := utils.FileToSlice("inputs/day4.txt")
	if err != nil {
		fmt.Println(err)
	}

	total := 0
	for _, line := range file {
		fmt.Println(line)
		game := utils.SliceMaker(line, ":")[1]

		result := cardHandler(game)
		total += result
	}

	fmt.Println(total)
}

func cardHandler(line string) int {
	parts := utils.SliceMaker(line, "|")
	winnersStr := utils.SliceMaker(parts[0], " ")
	myNumsStr := utils.SliceMaker(parts[1], " ")

	winners := toIntSlice(winnersStr)
	myNums := toIntSlice(myNumsStr)

	fmt.Println("winners: ", winners)
	fmt.Println("myNums: ", myNums)

	count := 0

	for _, num := range myNums {
		for _, win := range winners {
			if num == win {
				fmt.Printf("%d ", num)
				count++
			}
		}
	}

	fmt.Println("\nwinner count: ", count)

	return int(math.Pow(2, float64(count-1)))
}

func toIntSlice(str []string) []int {
	result := []int{}
	for _, num := range str {
		if num == "" {
			continue
		}

		num = strings.Trim(num, " ")

		intNum, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("error converting to integer: ", err)
			continue
		}

		result = append(result, intNum)
	}

	return result
}
