package util

import (
	"bufio"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func StringsArrayToInts(strings []string) ([]int, error) {
	var ints []int

	for _, i := range strings {
		j, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		ints = append(ints, j)
	}

	return ints, nil
}

func StringtoInt(s string) int {
	value, err := strconv.Atoi(s)
	Check(err)
	return value
}
