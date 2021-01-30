package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Template method using structs and interfaces
	fmt.Println("####################################################################")
	fmt.Println("------------Template Method Using Structs and Interfaces------------")
	gameOfChess := NewGameOfChess(0, 19, 0)
	PlayGame(gameOfChess)

	//Functional Template method (no structs or interfaces, just functions depending on functions)
	fmt.Println("####################################################################")
	fmt.Println("---------------Function Template Method Example---------------------")
	turn, maxTurns, currentPlayer := 0, 19, 0

	start := func() {
		fmt.Println("Starting the game of chess")
	}

	takeTurn := func() {
		fmt.Println("Turn", turn, "take by player", currentPlayer)
		turn++
		currentPlayer = 1 - currentPlayer
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGameFunctional(start, takeTurn, haveWinner, winningPlayer)

}

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

func PlayGame(game Game) {
	game.Start()

	for !game.HaveWinner() {
		game.TakeTurn()
	}

	fmt.Println("Winner is player: ", game.WinningPlayer())
}

type Chess struct {
	turn, maxTurns, currentPlayer int
}

func NewGameOfChess(turn, maxTurns, currentPlayer int) *Chess {
	return &Chess{turn, maxTurns, currentPlayer}
}

func (chess *Chess) Start() {
	fmt.Println("Starting new game of chess")
}

func (chess *Chess) TakeTurn() {
	chess.turn++
	fmt.Println("Turn " + strconv.Itoa(chess.turn) + " taken by player " + strconv.Itoa(chess.currentPlayer))
	chess.currentPlayer = 1 - chess.currentPlayer
}

func (chess *Chess) HaveWinner() bool {
	return chess.turn == chess.maxTurns
}

func (chess *Chess) WinningPlayer() int {
	return chess.currentPlayer
}

//Function Template Method

func PlayGameFunctional(start func(), takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	start()

	for !haveWinner() {
		takeTurn()
	}

	fmt.Println("Winner is player: ", winningPlayer())
}
