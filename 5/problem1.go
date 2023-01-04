package main

import "os"
import "fmt"

func parseStack(contents []uint8) [][]uint8 {
	/* Returns a slice of columns, each being a queue of the column's contents */
	
	var ret [][]uint8 /* Return value */
	contParse := true /* Whether or not to continue parsing */
	foundCrate := true /* Whether or not a crate was found in the current line */
	lines := [][]uint8{{0}} /* Parsed lines */
	numOfCols := 0 /* Find the number of columns */
	
	for contParse {
		linePos := 0 /* Index of the first ch of the current line in "contents" */
		lineNum := 0 /* Index of the current line */
		for i, ch := range contents {
			if (ch == '\n') {
				if !foundCrate {
					contParse = false
					lines = lines[:len(lines) - 1] /* Remove the last element. It'll be empty because it was parsing the column numbers */
					break
				}
				foundCrate = false
				linePos = i + 1
				lines = append(lines, make([]uint8, 0))
				lineNum++
			} else if ( ('A' <= ch) && (ch <= 'Z') ) {
				foundCrate = true
				columnNum := (i - linePos)/4

				/* If this is the highest index column we've found, save the index in numofcols */
				if (columnNum > numOfCols) {
					numOfCols = columnNum
				}
				

				/* If the current "line" array isn't long enough, keep appending 0s until it is */
				for (len(lines[lineNum]) <= columnNum) {
					lines[lineNum] = append(lines[lineNum], 0)
				}

				lines[lineNum][(i - linePos)/4] = ch
			}
		}
	}


	for i := 0; i <= numOfCols; i++ {
		ret = append(ret, make([]uint8, 0))
	}

	/* Rotate table (kind of) */
	for _, line := range lines {
		for col, val := range line {
			if (val != 0) {
				ret[col] = append(ret[col], val)
			}
		}
	}

	return ret

}

func parseInstructions(contents []uint8) [][3]uint8 {
	var ret [][3]uint8 /* Return value */
	instStart := 0
	stage := 0 /* 0=move, 1=from, 2=to */
	var tmp uint8 = 0
	parseNum := false /* Whether or not we are currently evaluating a number */
	retIndex := 0

	/* Just find where to start */
	for i, ch := range contents {
		if (ch == 'm') {
			if (contents[i + 1] == 'o' &&
				contents[i + 2] == 'v' &&
				contents[i + 3] == 'e') {
				instStart = i
				break
			}
		}
	}

	/* Actually parse instructions */
	for i := instStart; i < len(contents); i++ {
		ch := contents[i]
		if ('0' <= ch && ch <= '9') {
			parseNum = true
			tmp = (tmp * 10) + (ch - '0')
		} else {
			if parseNum {
				parseNum = false

				if (stage == 0) {
					ret = append(ret, [3]uint8{0, 0, 0})
				}
				
				ret[retIndex][stage] = tmp
				tmp = 0
				
				if (stage == 2) {
					retIndex++
					stage = 0
				} else {
					stage++
				}				
			}
		}
	}

	return ret
}

func main() {
	/* Parse input file */
	input, err := os.ReadFile("input.txt")
	if err != nil {panic(err)}

	stacks := parseStack(input)
	instructions := parseInstructions(input)
	
	for _, inst := range instructions {
		count, from, to := inst[0], inst[1] - 1, inst[2] - 1
		var slice []uint8

		if (count > uint8(len(stacks[from]))) {
			count = uint8(len(stacks[from]))
		}

		fmt.Println("From before:")
		fmt.Println(stacks[from])
		fmt.Println("To before:")
		fmt.Println(stacks[to])

		fmt.Printf("Move %d from %d to %d\n", count, from, to)
		/* Effectively "remove" a slice from the "From" stack */
		slice, stacks[from] = stacks[from][:count], stacks[from][count:]

		stacks[to] = append(slice, stacks[to]...)

		fmt.Println("From after:")
		fmt.Println(stacks[from])
		fmt.Println("To after:")
		fmt.Println(stacks[to])
		fmt.Println()
		
	}

	fmt.Println(stacks)
}
