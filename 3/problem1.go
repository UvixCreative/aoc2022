package main

import "os"
import "fmt"
import "strings"

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		first := line[0:len(line)/2]
		second := line[(len(line)/2):len(line)]

		fmt.Printf("Whole: %s\nFirst: %s\nSecond: %s\n\n", line, first, second)
	}
}
