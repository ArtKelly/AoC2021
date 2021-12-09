package challenges

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

func Day6(filepath string) {
	lines := util.ReadLines(filepath)
	values := make([]int, len(strings.Split(lines[0], ",")))
	state := make([]int, 9)
	for i, v := range strings.Split(lines[0], ",") {
		values[i], _ = strconv.Atoi(v)
		state[values[i]]++
	}
	for i := 0; i < 80; i++ {
		shift(&state)
	}
	fmt.Printf("Part 1: %d\n", sumRow(state))
	for i := 0; i < 176; i++ {
		shift(&state)
	}
	fmt.Printf("Part 2: %d\n", sumRow(state))
}

func shift(state *[]int) {
	newFish := (*state)[0]
	(*state)[7] += newFish
	(*state) = (*state)[1:]
	(*state) = append(*state, newFish)
}
