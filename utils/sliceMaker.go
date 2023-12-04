package utils

import "strings"

func SliceMaker(str string, seperator string) []string {
	slice := strings.Split(str, seperator)

	for i := 0; i < len(slice); i++ {
		slice[i] = strings.Trim(slice[i], " ") // remove leading space
	}

	return slice
}
