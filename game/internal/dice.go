package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once
var pickDice int

func rollDice() {
	dice := []int{1, 2, 3, 4, 5, 6}
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	pickDice = dice[rand.Intn(len(dice))]
	fmt.Println("pickDice:", pickDice)
}

func guessDice(text string) (bool, string, string) {
	if text != "small" && text != "big" {
		panic("You can choose only one big or small")
	}
	if text == "small" {
		if pickDice == 1 || pickDice == 2 || pickDice == 3 {
			rollDice()
			return true, text, "win"
		}
		rollDice()
		return false, text, "lost"
	}
	if text == "big" {
		if pickDice == 4 || pickDice == 5 || pickDice == 6 {
			rollDice()
			return true, text, "win"
		}
		rollDice()
		return false, text, "lost"

	}
	return false, "", ""
}

func stopRollingDice() {
	pickDice = 0
}
