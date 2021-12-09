package challenges

import (
	"fmt"

	util "github.com/ArtKelly/AoC2021/Util"
)

// TODO: This is bad, pls fix in future
func Day5(filepath string) {
	lines := util.ReadLines(filepath)
	var seaFloor [1000][1000]int
	for _, l := range lines {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(l, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		util.Check(err)
		if x1 == x2 {
			// Handle vertical
			if y2 < y1 {
				tmp := y1
				y1 = y2
				y2 = tmp
			}
			for i := y1; i <= y2; i++ {
				(seaFloor)[i][x1]++
			}
		} else if y1 == y2 {
			// Handle horizontal
			if x2 < x1 {
				tmp := x1
				x1 = x2
				x2 = tmp
			}
			for i := x1; i <= x2; i++ {
				(seaFloor)[y1][i]++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", sumVents(seaFloor))
	var newSeaFloor [1000][1000]int
	for _, l := range lines {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(l, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		util.Check(err)
		if x1 == x2 {
			// Handle vertical
			if y2 < y1 {
				tmp := y1
				y1 = y2
				y2 = tmp
			}
			for i := y1; i <= y2; i++ {
				(newSeaFloor)[i][x1]++
			}
		} else if y1 == y2 {
			// Handle horizontal
			if x2 < x1 {
				tmp := x1
				x1 = x2
				x2 = tmp
			}
			for i := x1; i <= x2; i++ {
				(newSeaFloor)[y1][i]++
			}
		} else {
			// handle diag
			delta := (y2 - y1) / (x2 - x1)
			if delta == 1 {
				if x2 < x1 {
					tmp1, tmp2 := y1, x1
					y1, x1 = y2, x2
					y2, x2 = tmp1, tmp2
				}
				for i, j := x1, y1; i <= x2 && j <= y2; i, j = i+1, j+1 {
					newSeaFloor[j][i]++
				}
			} else {
				if y2 < y1 {
					tmp1, tmp2 := y1, x1
					y1, x1 = y2, x2
					y2, x2 = tmp1, tmp2
				}
				for i, j := x1, y1; i >= x2 && j <= y2; i, j = i-1, j+1 {
					newSeaFloor[j][i]++
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", sumVents(newSeaFloor))
}

func sumVents(seaFloor [1000][1000]int) int {
	sum := 0
	for _, i := range seaFloor {
		for _, j := range i {
			if j > 1 {
				sum++
			}
		}
	}
	return sum
}
