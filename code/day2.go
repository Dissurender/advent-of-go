package code

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MAP = map[string]int

var validData = MAP{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Day2() {
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		fmt.Printf("error reading input file: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	gamePower := 0

	for {
		line, err := reader.ReadString('\n') // read line from txt file

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			return
		}

		round, values := parse(line)

		validVal, power := validateValues(values, validData)
		fmt.Println("added power: ", power)
		gamePower = gamePower + power

		if !validVal {
			fmt.Println("skipped round ", round)
			continue
		}

		sum = sum + round

		fmt.Println()

	}

	fmt.Printf("The final un-impossible game sum-counterino is: %d\n", sum)
	fmt.Printf("It's power lever is exactly %d thousand!\n", gamePower)
}

func parse(str string) (int, []MAP) {
	parts := sliceMaker(str, ":")

	game := gameRound(parts[0])

	round := roundsHandler(parts[1])
	rounds := buildRounds(round)

	return game, rounds
}

func gameRound(round string) int {
	sep := sliceMaker(round, " ") // ["Game", #]

	roundInt, err := strconv.Atoi(sep[1])
	if err != nil {
		return 0
	}

	return roundInt
}

func roundsHandler(round string) []string {
	parts := sliceMaker(round, ";") // separate pulls
	return parts
}

func buildRounds(rounds []string) []MAP {
	var result []MAP

	for _, round := range rounds {
		hands := sliceMaker(round, ",")

		handMap := buildHands(hands)
		result = append(result, handMap)
	}

	return result
}

func buildHands(hands []string) MAP {
	result := make(MAP)

	for _, hand := range hands {
		parts := sliceMaker(hand, " ")
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error parsing count: ", err)
			continue
		}

		color := parts[1]
		result[color] = count
	}

	return result
}

func validateValues(vals []MAP, validData MAP) (bool, int) {
	power := 1
	valideGame := true
	minimums := MAP{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, hands := range vals {
		for color, count := range hands {

			color = strings.Trim(color, "\n")

			if count > validData[color] {
				fmt.Printf("failed ==== color: '%s' count: '%d'\n", color, count)
				valideGame = false
			}
			if count > minimums[color] {
				minimums[color] = count
			}
		}
	}

	power = minimums["red"] * minimums["green"] * minimums["blue"]

	return valideGame, power
}

func sliceMaker(str string, seperator string) []string {
	slice := strings.Split(str, seperator)

	for i := 0; i < len(slice); i++ {
		slice[i] = strings.Trim(slice[i], " ") // remove leading space
	}

	return slice
}
