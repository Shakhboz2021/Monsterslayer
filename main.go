package main

import (
	"fmt"

	"github.com/Shakhboz2021/Monsterslayer/interaction"
)

var currentRound = 0

func main() {

	startGame()

	winner := "" // "Player" || "Monster"
	for winner == "" {
		winner = executeRound()
	}
	endGame()

}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	fmt.Println(userChoice)
	return ""
}

func endGame() {

}
