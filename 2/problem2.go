package main

import "os"
import "fmt"

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	var myScore int
	var myWins int

	/* First dimension is opponent shape (0 = R, 1 = P, 2 = S)
	   Second dimension gives value to attain desired status (0 = lose,
	   1 = draw, 2 = win) */
	winMatrix := [3][3]int{{3, 1, 2}, {1, 2, 3}, {2, 3, 1}}
	
	for i := 0; i < len(input); i += 4 {
		var opponentPick int
		var myPick int
		var desiredStatus int
		var pointsWon int

		/* Determine desired status based on second column */
		switch input[i + 2] {
		case 'X':
			desiredStatus = 0
		case 'Y':
			desiredStatus = 1
		case 'Z':
			desiredStatus = 2
		}

		/* Convert pick to comparable numeric value */
		switch input[i] {
		case 'A':
			opponentPick = 1
		case 'B':
			opponentPick = 2
		case 'C':
			opponentPick = 3
		}

		myPick = winMatrix[opponentPick - 1][desiredStatus]
		
		/* Calculate difference */
		result := opponentPick - myPick

		fmt.Println(result)

		switch result {
		/* Win */
		case -1:
			fallthrough
		case 2:
			myWins++
			pointsWon = 6
		/* Tie */
		case 0:
			pointsWon = 3
		/* Loss */
		case 1:
			fallthrough
		case -2:
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
