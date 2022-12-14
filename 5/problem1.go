package main

import "os"
import "fmt"

func parse(contents []uint8) [][]uint8 {
	var ret [][]uint8
	contParse := true
	foundCrate := false
	
	for contParse {
		linePos := 0
		lineNum := 0
		for i := 2; i <= len(
		for i, ch := range contents {
			if (ch == '\n') {
				if !foundCrate {
					return ret
				}
				fmt.Println(ret[lineNum])
				foundCrate = false
				linePos = i + 1
				lineNum++
			} else if ( ('A' <= ch) && (ch <= 'Z') ) {
				foundCrate = true

				ret[lineNum][(i - linePos)/4] = ch
				fmt.Printf("%c in column %d\n", ch, (i - linePos)/4)
			}
		}
	}

	return ret

}

func main() {
	/* Parse input file */
	input, err := os.ReadFile("input.txt")
	if err != nil {panic(err)}
	
	fmt.Println(parse(input))
}
