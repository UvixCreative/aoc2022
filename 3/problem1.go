package main

import "os"
import "fmt"
import "strings"

func findSharedChar(first string, second string) (rune, bool) {
	/* Things for building a set */
	type void struct{}
	var member void

	/* Function variables */
	secondSet := make(map[rune]void)	
	var foundCh rune = 0
	var found bool = false

	/* Fill the set with each character from second string */
	for _, ch := range second {
		secondSet[ch] = member
	}

	/* See if any character from string 1 is in secondSet */	
	for _, ch := range first {
		if _, ok := secondSet[ch]; ok {
			found = true
			foundCh = ch
			break
		}
	}

	return foundCh, found
}

func main() {
	/* Parse input file into lines */
	input, err := os.ReadFile("input.txt")
	if err != nil {panic(err)}
	lines := strings.Split(string(input), "\n")

	prioritySum := 0
	for _, line := range lines {
		first := line[:len(line)/2]
		second := line[(len(line)/2):]

		ch, found := findSharedChar(first, second)

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
