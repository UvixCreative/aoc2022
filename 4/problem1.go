package main

import "os"
import "fmt"

func parse(contents []uint8) [][4]int {
	/* Takes the input file and parses the characters
	   to return a slice of int arrays. Each int array
	   represents the start and end number of the ranges
	   for each elf */
	
	pairNum := 0 /* Pair number/index */
	i := 0 /* Used for index within a single pair */

	tmp := 0
	tmpArr := [4]int{0}

	var elfRanges [][4]int /* Slice of range pairs */

	for _, ch := range contents {
		if (ch == ',' || ch == '-' || ch =='\n') {
			tmpArr[i] = tmp
			tmp = 0
			i++
		} else if ('0' <= ch || ch <= '9') {
			ch -= '0'
			tmp = (tmp * 10) + int(ch)
		}

		if (ch == '\n') {
			i = 0
			elfRanges = append(elfRanges, tmpArr)
			pairNum++
		}
	}

	return elfRanges
}

func isBetween(comp int, min int, max int) bool {
	return (min <= comp && comp <= max)
}

func isSubset(comp []int, compRange []int) bool {
	return (isBetween(comp[0], compRange[0], compRange[1]) &&
		isBetween(comp[1], compRange[0], compRange[1]))
	
}

func main() {
	/* Parse input file */
	input, err := os.ReadFile("input.txt")
	if err != nil {panic(err)}

	ret := 0

	for _, pair := range parse(input) {
		if (isSubset(pair[0:2], pair[2:4]) || isSubset(pair[2:4], pair[0:2])) {
			ret++
		}
	}

	fmt.Printf("%d elf pairs have a complete overlap\n", ret)
}
