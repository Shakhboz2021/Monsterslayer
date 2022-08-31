package main

import (
	"github.com/Shakhboz2021/Monsterslayer/actions"
	"github.com/Shakhboz2021/Monsterslayer/interaction"
)

var currentRound = 0

func main() {

	startGame()

	winner := "" // "Player" || "Monster"
	for winner == "" {
		winner = executeRound()
	}
	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDamage int
	var playerHealValue int
	var monsterAttackDamage int

	if userChoice == "ATTACK" {
		playerAttackDamage = actions.AttackingToMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		monsterAttackDamage = actions.AttackingToMonster(isSpecialRound)
	}

	actions.AttackingToPlayer()

	userHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:              userChoice,
		PlayerHealth:        userHealth,
		MonsterHealth:       monsterHealth,
		PlayerAttackDamage:  playerAttackDamage,
		PlayerHealValue:     playerHealValue,
		MonsterAttackDamage: monsterAttackDamage,
	}

	interaction.PrintRoundStatistics(&roundData)

	if userHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}
	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
}
