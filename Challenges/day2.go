package challenges

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

func Day2() {
	inputs, horizontal, depth := util.ReadLines("Inputs/day2.txt"), 0, 0
	for _, command := range inputs {
		command, value := parseCommandString(command)
		switch command {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	fmt.Printf("Part 1: %d\n", (depth * horizontal))

	horizontal, depth, aim := 0, 0, 0
	for _, command := range inputs {
		command, value := parseCommandString(command)
		switch command {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	fmt.Printf("Part 2: %d\n", (depth * horizontal))
}

func parseCommandString(s string) (string, int) {
	commands := strings.Fields(s)
	i, _ := strconv.Atoi(commands[1])
	return commands[0], i
}
