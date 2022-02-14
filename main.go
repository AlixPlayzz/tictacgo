package main

import (
	"bufio"
	"fmt"
	"time"

	"os"
	"strings"
)

// random vars
var whosTurn string = "X" // X always has the first turn

// main array[] (byte array) with the game in it
var game [9]string = [9]string{"-", "-", "-",
	"-", "-", "-",
	"-", "-", "-"}

func main() {

	// loops until game ends
	for i := 0; i < 9; i++ {

		// checks for if "X" has won
		if game[0] == "X" && game[1] == "X" && game[2] == "X" ||
			game[0] == "X" && game[3] == "X" && game[6] == "X" ||
			game[6] == "X" && game[7] == "X" && game[8] == "X" ||
			game[2] == "X" && game[5] == "X" && game[8] == "X" ||
			game[0] == "X" && game[4] == "X" && game[8] == "X" ||
			game[2] == "X" && game[4] == "X" && game[6] == "X" ||
			game[1] == "X" && game[4] == "X" && game[7] == "X" ||
			game[3] == "X" && game[4] == "X" && game[5] == "X" {

			xWon()
			return
		}

		// checks for if "O" has won
		if game[0] == "O" && game[1] == "O" && game[2] == "O" ||
			game[0] == "O" && game[3] == "O" && game[6] == "O" ||
			game[6] == "O" && game[7] == "O" && game[8] == "O" ||
			game[2] == "O" && game[5] == "O" && game[8] == "O" ||
			game[0] == "O" && game[4] == "O" && game[8] == "O" ||
			game[2] == "O" && game[4] == "O" && game[6] == "O" ||
			game[1] == "O" && game[4] == "O" && game[7] == "O" ||
			game[3] == "O" && game[4] == "O" && game[5] == "O" {

			oWon()
			return
		}

		round()
	}

}

// called each round
func round() {

	gameOutput()

	// all three lines of values
	var values []string

	for i := 1; i <= 3; i++ {
		// scanner
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		sT := s.Text()

		// need to put this check here so runtime doesnt panic
		if len(sT) != 3 {
			fmt.Println("You did something wrong. Try again.")
			time.Sleep(2 * time.Second)

			// restart the round
			round()
			return
		}

		// makes an array with the users input line
		sTSlice := strings.Split(sT, "")
		values = append(values, strings.ToUpper(sTSlice[0]), strings.ToUpper(sTSlice[1]), strings.ToUpper(sTSlice[2]))
	}

	// convert slice array into string
	valuesString := strings.Join(values, "")

	// does a bunch of checks to make sure user has done correctly
	if strings.Contains(valuesString, "X") && strings.Contains(valuesString, "O") ||
		!strings.Contains(valuesString, "X") && !strings.Contains(valuesString, "O") ||
		!strings.Contains(valuesString, "-") ||
		len(valuesString) != 9 ||
		!strings.Contains(valuesString, whosTurn) {

		fmt.Println("You did something wrong. Try again.")
		time.Sleep(2 * time.Second)

		// restart the round
		round()
		return
	}

	// combine this rounds move with the "game" var
	for i := 0; i < 9; i++ {

		// checks that X or O doesnt overlap with existing game variable
		if game[i] == "-" && values[i] != "-" {
			game[i] = values[i]

			// who turns variable update
			if values[i] == "X" {
				whosTurn = "O"
			}
			if values[i] == "O" {
				whosTurn = "X"
			}

			break
		}
	}
}

// output how the game is looking
func gameOutput() {
	for i := 0; i < 100; i++ {
		fmt.Println("")
	}
	fmt.Println("Game:")
	fmt.Println(game[0] + game[1] + game[2])
	fmt.Println(game[3] + game[4] + game[5])
	fmt.Println(game[6] + game[7] + game[8])
	fmt.Println("It's now \"" + string(whosTurn) + "\"'s turn")
}

func xWon() {
	gameOutput()
	fmt.Println("X won the game, and O has a skill issue! lol!!!!!")
	time.Sleep(1 * time.Second)
}
func oWon() {
	gameOutput()
	fmt.Println("O won the game, and X has a skill issue! lol!!!!!")
	time.Sleep(1 * time.Second)
}
