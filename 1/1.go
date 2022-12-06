package main

import "os"
import "fmt"
import "reflect"
import "strings"
import "strconv"

func main() {
	input, err := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	var elve_cals [5]int

	if err != nil {
		panic(err)
	}

	for _, s := range lines {
		if (s == "") {
			fmt.Println("EMPTY")
		} else {
			fmt.Println(s)
			fmt.Println(strconv.Atoi(s))
		}
	}
	// fmt.Print(lines)
	fmt.Printf("%s", reflect.TypeOf(lines))
}
