package challenges

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

func Day4(filepath string) {
	var (
		boards [][][]int
		board  [][]int
	)

	lines := util.ReadLines(filepath)
	for i, l := range lines {
		if i <= 1 {
			// Skip first line
			continue
		}
		// blank line indicates end of previous board
		if l == "" {
			boards = append(boards, board)
			board = make([][]int, 0)
			continue
		}

		nums := strings.Fields(l)
		rows := make([]int, 0)
		for _, number := range nums {
			i, _ := strconv.Atoi(number)
			rows = append(rows, i)
		}
		board = append(board, rows)
	}
	boards = append(boards, board)

	// Loop through all the chosen bingo numbers
	noWon := 0
	for _, i := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(i)
		for _, b := range boards {
			if !boardHasWon(b) {
				updateBoard(&b, i)
				if boardHasWon(b) {
					noWon++
					if noWon == 1 {
						s := sumBoard(b)
						fmt.Printf("Part1: %d\n", s*i)
					} else if noWon == len(boards) {
						s := sumBoard(b)
						fmt.Printf("Part2: %d\n", s*i)
					}
				}
			}
		}
	}
}

func boardHasWon(board [][]int) bool {
	for _, row := range board {
		if sumRow(row) == 0 {
			return true
		}
	}
	tmpBoard := transposeInt(board)
	for _, row := range tmpBoard {
		if sumRow(row) == 0 {
			return true
		}
	}
	return false
}

func updateBoard(board *[][]int, chosenNum int) {
	for i, row := range *board {
		for j, num := range row {
			if num == chosenNum {
				(*board)[i][j] = 0
			}
		}
	}
}

func sumBoard(board [][]int) int {
	sum := 0
	for _, row := range board {
		sum += sumRow(row)
	}
	return sum
}

func sumRow(row []int) int {
	sum := 0
	for _, i := range row {
		sum += i
	}
	return sum
}

func transposeInt(slice [][]int) [][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
