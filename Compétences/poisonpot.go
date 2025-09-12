package main

import (
	"fmt"
	"time"
)

func poisonPot(currentHP, maxHP int) int {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)

		currentHP -= 10
		if currentHP < 0 {
			currentHP = 0
		}

		fmt.Printf("PV: %d / %d\n", currentHP, maxHP)
	}
	return currentHP
}

func main() {
	currentHP := 50
	maxHP := 100

	finalHP := poisonPot(currentHP, maxHP)
	fmt.Printf("PV finaux après le poison : %d\n", finalHP)
}
