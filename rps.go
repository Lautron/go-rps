package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	rock = iota
	paper
	scissors
)

const (
	win = iota
	tie
	lose
)

type ResultMatrix [3][3]int
type MoveArray [3]string
type MsgArray [3]string

// maybe change name to eventNames
func getMoveNames() MoveArray {
	return [3]string{"Rock", "Paper", "Scissors"}
}
func getResultMsgs() MsgArray {
	return [3]string{"You Win!", "There was a Tie!", "You Lose!"}
}

func getResults() ResultMatrix {
	var results ResultMatrix

	results[rock][rock] = tie
	results[rock][paper] = lose
	results[rock][scissors] = win

	results[paper][rock] = win
	results[paper][paper] = tie
	results[paper][scissors] = lose

	results[scissors][rock] = lose
	results[scissors][paper] = win
	results[scissors][scissors] = tie

	return results
}

func getAiMove() int {
	return rand.Intn(3)
}

func getUserInput() (string, error) {
	var r string
	_, err := fmt.Scanf("%s", &r)

	if err != nil {
		return "", err
	}

	return r, nil
}

func getUserMove() (int, error) {
	r, err := getUserInput()
	if err != nil {
		return 0, err
	}
	result, _ := strconv.Atoi(r)

	return result - 1, nil
}

func main() {
	results := getResults()
	resultMsgs := getResultMsgs()
	moveNames := getMoveNames()

	for {
		aiMove := getAiMove()
		fmt.Print("Choose a move:\n1) Rock\n2) Paper\n3) Scissors\nYour move: ")
		userMove, moveErr := getUserMove()
		fmt.Println()
		if moveErr != nil {
			fmt.Println("Invalid move! Try again")
			continue
		}

		fmt.Printf("%s vs %s\n\n", moveNames[userMove], moveNames[aiMove])
		result := results[userMove][aiMove]

		fmt.Println(resultMsgs[result])
		fmt.Println("Play again? (y/n)")
		userAnswer, moveErr := getUserInput()
		if moveErr != nil {
			fmt.Println("Invalid answer! Try again")
			continue
		}

		if userAnswer == "n" {
			break
		}
	}
}
