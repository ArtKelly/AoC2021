package challenges

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

func Day7(filepath string) {
	lines := util.ReadLines(filepath)
	positions := strings.Split(lines[0], ",")
	intPositions := make([]int, len(positions))
	for i, v := range positions {
		intPositions[i], _ = strconv.Atoi(v)
	}
	fuel := 0
	medianValue := getMedian(intPositions)
	for _, v := range intPositions {
		diff := int(math.Abs(float64(v) - float64(medianValue)))
		fuel += diff
	}
	fmt.Printf("Part 1: %d\n", fuel)

	fuel = 0
	avgValue := getAvg(intPositions)
	for _, v := range intPositions {
		diff := int(math.Abs(float64(v) - float64(avgValue)))
		fuel += (diff * (diff + 1)) / 2
	}
	fmt.Printf("Part 2: %d\n", fuel)
}

func getMedian(positions []int) int {
	sort.Ints(positions)
	middle := len(positions) / 2
	medianValue := 0
	if len(positions)%2 == 1 {
		medianValue = positions[middle]
	} else {
		medianValue = (positions[middle-1] + positions[middle]) / 2
	}
	return medianValue
}

func getAvg(positions []int) int {
	sum := sumRow(positions)
	return int(sum / len(positions))
}
