package main

import "os"
import "fmt"
import "strings"
import "strconv"

func quicksort(arr []int) []int {
	if (len(arr) <= 1) {
		return arr
	}
	pivot := (len(arr) / 2)

	pivotval := arr[pivot]

	var lesser []int
	var greater []int
	var eq []int
	
	for _, v := range arr {
		if (v < pivotval) {
			lesser = append(lesser, v)
		} else if (v > pivotval) {
			greater = append(greater, v)
		} else {
			eq = append(eq, v)
		}
	}

	lesser = quicksort(lesser)
	greater = quicksort(greater)
	eq = append(lesser, eq...)

	return append(eq, greater...)
}

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
			// fmt.Printf("Elf %d: %d calories\n", i, tmp_total)
			elve_cals = append(elve_cals, tmp_total)
			tmp_total = 0
			i++
		} else {
			val, _ := strconv.Atoi(s)
			tmp_total += val
		}
	}

	elve_cals = quicksort(elve_cals)
	elve_cals_len := len(elve_cals)


	fmt.Printf("\nThe elf with the most calories is carrying %d calories\n", elve_cals[elve_cals_len - 1] )

	var topthree int
	for i:= 3; i > 0; i-- {
		topthree += elve_cals[elve_cals_len - i]
	}

	
	fmt.Printf("\nThe top 3 elves have %d combined calories\n", topthree)
}
