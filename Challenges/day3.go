package challenges

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

func Day3(filepath string) {
	inputs := util.ReadLines(filepath)
	new := transpose(inputs)
	gamma, epsilon := "", ""
	for _, line := range new {
		bit := findCommonBit(line)
		if bit == '1' {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaNum, _ := strconv.ParseUint(gamma, 2, 64)
	epsilonNum, _ := strconv.ParseUint(epsilon, 2, 64)
	fmt.Printf("Part1: %d\n", gammaNum*epsilonNum)

	// Part2
	oxyRating, _ := strconv.ParseUint(caluclateRating(inputs, 0, true), 2, 64)
	c02Rating, _ := strconv.ParseUint(caluclateRating(inputs, 0, false), 2, 64)
	fmt.Printf("Part2: %d\n", oxyRating*c02Rating)

}

func caluclateRating(slice []string, index int, commonBit bool) string {
	// Base case, break if slice has one element
	if len(slice) == 1 {
		return slice[0]
	}

	oxySlice, co2Slice := make([]string, 0), make([]string, 0)

	for _, line := range slice {
		if rune(line[index]) == '1' {
			oxySlice = append(oxySlice, line)
		} else {
			co2Slice = append(co2Slice, line)
		}
	}

	if len(oxySlice) >= len(co2Slice) == commonBit {
		return caluclateRating(oxySlice, index+1, commonBit)
	}
	return caluclateRating(co2Slice, index+1, commonBit)
}

func transpose(slice []string) []string {
	length := len(slice)
	newArray := make([]string, len(slice[0]))
	for i := 0; i < len(slice[0]); i++ {
		newStr := ""
		for j := 0; j < length; j++ {
			newStr += string([]rune(slice[j])[i])
		}
		newArray[i] = newStr
	}
	return newArray
}

func findCommonBit(s string) rune {
	ones := strings.Count(s, "1")
	zeroes := strings.Count(s, "0")
	if ones >= zeroes {
		return '1'
	}
	return '0'
}
