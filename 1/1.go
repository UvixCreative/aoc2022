package main

import "os"
import "fmt"
import "strings"
import "strconv"

func main() {
	input, err := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	var elve_cals []int
	var tmp_total int
	i := 0

	if err != nil {
		panic(err)
	}

	for _, s := range lines {
		if (s == "") {
			fmt.Printf("Elf %d: %d calories\n", i, tmp_total)
			elve_cals = append(elve_cals, tmp_total)
			tmp_total = 0
			i++
		} else {
			val, _ := strconv.Atoi(s)
			tmp_total += val
		}
	}

	var biggest int
	var biggestIdx int
	for index, val := range elve_cals {
		if (val > biggest) {
			biggest = val
			biggestIdx = index
		}
	}

	fmt.Printf("\nElf %d has the greatest, at %d calories\n", biggestIdx, biggest)
}
