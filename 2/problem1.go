package main

import "os"
import "fmt"

func main() {
	input, err := os.ReadFile("input.txt")

	var myScore int
	var myWins int
	
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(input); i += 4 {
		opponentPick := input[i]
		myPick := input[i + 2]
		var pointsWon int

		/* Convert pick to comparable numeric value */
		switch opponentPick {
		case 'A':
			opponentPick = 1
		case 'B':
			opponentPick = 2
		case 'C':
			opponentPick = 3
		}

		switch myPick {
		case 'X':
			myPick = 1
		case 'Y':
			myPick = 2
		case 'Z':
			myPick = 3
		}
		
		/* Calculate difference */
		result := opponentPick - myPick

		switch int(result) {
		/* Win */
		case -1:
			fallthrough
		case -2:
			myWins++
			pointsWon = 6
		/* Tie */
		case 0:
			pointsWon = 3
		/* Loss */
		case 1:
			fallthrough
		case 2:
			pointsWon = 0
		}

		/* Add score based on which shape I picked */
		matchPoints := int(myPick) + pointsWon

		/* Add match points */
		myScore += matchPoints

		fmt.Println(opponentPick, myPick)
		fmt.Printf("Points won: %d (Current total: %d)\n", matchPoints, myScore)
	}
}
