package challenges

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	util "github.com/ArtKelly/AoC2021/Util"
)

var reference []string = []string{
	"abcefg",
	"cf",
	"acdeg",
	"acdfg",
	"bcdf",
	"abdfg",
	"abdefg",
	"acf",
	"abcdefg",
	"abcdfg",
}

// fmt.Printf("%s %s %s %s", SortString("cg"), SortString("cgb"), SortString("egdcabf"), SortString("gecf"))
func Day8(filepath string) {

	lines := util.ReadLines(filepath)
	sum, sum2 := 0, 0
	seenCodes := make(map[string]int)
	for _, v := range lines {
		codes := strings.Split(v, " | ")
		firstCodes := strings.Split(codes[0], " ")
		secondCodes := strings.Split(codes[1], " ")
		allCodes := append(firstCodes, secondCodes...)
		sum2 += decode(allCodes, secondCodes)
		for _, x := range secondCodes {
			seenCodes[x] += 1
			if len(x) == 2 || len(x) == 3 || len(x) == 4 || len(x) == 7 {
				sum++
			}
		}

	}
	fmt.Printf("Part 1: %d", sum)
	fmt.Printf("Part 2: %d", sum2)
}

// 0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....

//  5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

func decode(all []string, second []string) int {
	decoder := make(map[int]string)
	unknowns := []string{}
	decoder[8] = "abcdefg"
	for _, v := range all {
		if len(v) == 2 {
			decoder[1] = v
		} else if len(v) == 3 {
			decoder[7] = v
		} else if len(v) == 4 {
			decoder[4] = v
		} else if len(v) == 7 {
			decoder[8] = v
		} else {
			unknowns = append(unknowns, v)
		}
	}
	for _, v := range unknowns {
		if len(v) == 5 {
			if len(Intersection([]rune(v), []rune(decoder[7]))) == 3 {
				decoder[3] = v
				// fmt.Printf("%v\n %v\n", decoder, string(Intersection([]rune(v), []rune(decoder[7]))))
			} else if len(Intersection([]rune(v), []rune(decoder[4]))) == 2 {
				decoder[2] = v
			} else {
				decoder[5] = v
			}
			// Could be 2, 3 or 5
		} else if len(v) == 6 {
			if len(Intersection([]rune(v), []rune(decoder[7]))) == 2 {
				decoder[6] = v
			} else if len(Intersection([]rune(v), []rune(decoder[4]))) == 4 {
				decoder[9] = v
			} else {
				decoder[0] = v
			}
			// Could be 0, 6, 9
		}
	}
	sum := ""
	for _, v := range second {
		for key, val := range decoder {
			if SortString(v) == SortString(val) {
				sum += string(strconv.Itoa(key))
			}
		}
	}
	sum2, _ := strconv.Atoi(sum)
	fmt.Println(sum2)
	return sum2
}
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Intersection(a, b []rune) (c []rune) {
	m := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
