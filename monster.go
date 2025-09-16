package main

import "fmt"


func Monster() {
    regirock := Monster{
        Name:        "Regirock",
        MaxHP:       150,
        CurrentHP:   150,
        AttackPower: 30,
    }

    fmt.Printf("Nom: %s\n", regirock.Name)
    fmt.Printf("PV: %d/%d\n", regirock.CurrentHP, regirock.MaxHP)
    fmt.Printf("Attaque: %d\n", regirock.AttackPower)
}

func initGoblin() Monster {
	return Monster{
		Name: 		"Golem d'entraînement"
		MaxHP:		40
		CurrentHP:	40
		AttackPower:5
	}
}

func main() {
	golem := initGoblin()

	fmt.Printf("Nom %s\n", golem.Name)
	fmt.Printf("PV: %d/%d\n", golem.CurrentHP, golem.MaxHP)
	fmt.Printf("Attaque: %\n," golem.AttackPower)
}
