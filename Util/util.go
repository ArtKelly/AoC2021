package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadLines(path string) []string {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())
	return lines
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
