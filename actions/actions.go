package actions

import (
	"math/rand"
	"time"
)

var randomSource = rand.NewSource(time.Now().UnixNano())
var randomGenerator = rand.New(randomSource)

func AttackMonster() {

}

func HealPlayer() {

}

func AttackPlayer() {

}

func generateRandomBetween(min int, max int) int {
	return randomGenerator.Intn(max-min) + min
}
