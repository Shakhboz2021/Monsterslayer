package actions

import (
	"math/rand"
	"time"
)

var randomSource = rand.NewSource(time.Now().UnixNano())
var randomGenerator = rand.New(randomSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackingToMonster(isSpecialAttack bool) int {
	minAttack := PLAYER_ATTACK_MIN_DAMAGE
	maxAttack := PLAYER_ATTACK_MAX_DAMAGE
	if isSpecialAttack {
		minAttack = PLAYER_SPECIAL_ATTACK_MIN_DAMAGE
		maxAttack = PLAYER_SPECIAL_ATTACK_MAX_DAMAGE
	}
	dmgValue := generateRandomBetween(minAttack, maxAttack)
	currentMonsterHealth -= dmgValue

	return dmgValue
}

func HealPlayer() int {

	healValue := generateRandomBetween(PLAYER_HEAL_MIN_DAMAGE, PLAYER_HEAL_MAX_DAMAGE)
	if healValue+currentPlayerHealth > PLAYER_HEALTH {
		currentPlayerHealth = PLAYER_HEALTH
	} else {
		currentPlayerHealth += healValue
	}
	return healValue
}

func AttackingToPlayer() int {
	dmgValue := generateRandomBetween(MONSTER_ATTACK_MIN_DAMAGE, MONSTER_ATTACK_MAX_DAMAGE)
	currentPlayerHealth -= dmgValue
	return dmgValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandomBetween(min int, max int) int {
	return randomGenerator.Intn(max-min) + min
}
