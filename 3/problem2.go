package main

import "os"
import "fmt"
import "strings"

func findSharedChar(first string, second string, third string) (rune, bool) {
	/* Things for building a set */
	type void struct{}
	var member void

	/* Function variables */
	secondSet := make(map[rune]void)
	thirdSet := make(map[rune]void)
	var foundCh rune = 0
	var found bool = false

	/* Fill the set with each character from second string */
	for _, ch := range second {
		secondSet[ch] = member
	}

	/* Same thing for third string */
	for _, ch := range third {
		thirdSet[ch] = member
	}

	/* See if any character from string 1 is in secondSet */
	for _, ch := range first {
		if _, secOk := secondSet[ch]; secOk {
			if _, thirOk := thirdSet[ch]; thirOk {
				found = true
				foundCh = ch
				break
			}
		}
	}

	return foundCh, found
}

func main() {
	/* Parse input file into lines */
	input, err := os.ReadFile("input.txt")
	if err != nil {panic(err)}
	trimmed := strings.Trim(string(input), " ")
	lines := strings.Split(trimmed, "\n")

	prioritySum := 0
	for i := 0; i < (len(lines) - 3); i += 3 {
		first := lines[i]
		second := lines[i + 1]
		third := lines[i + 2]

		ch, found := findSharedChar(first, second, third)

		if found {
			priority := int(ch - 'A')
			upperOffset := 27
			lowerOffset := -31
			
			if (31 < priority && priority < 58) {
				/* Lowercase letters */				
				priority += lowerOffset
			} else if (0 <= priority && priority < 27) {
				/* Uppercase letters */				
				priority += upperOffset
			}

			prioritySum += priority
		}
	}

	fmt.Printf("Total priority: %d\n", prioritySum)
}
