package main

import (
	"fmt"
	"strings"
)

// Définition de la structure Character
type Character struct {
	Name      string
	Class     string
	Level     int
	HpMax     int
	HpActual  int
	Inventory []string
}

// Fonction qui initialise un personnage
func initCharacter() Character {
	var name, class string

	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)

	fmt.Print("Choisissez une classe (Chevalier, Archer, Magicien) : ")
	fmt.Scanln(&class)

	// Normalisation de la casse
	class = strings.Title(strings.ToLower(class))

	level := 1
	var hpMax, hpActual int

	switch class {
	case "Chevalier":
		hpMax = 150
		hpActual = 150
	case "Archer":
		hpMax = 75
		hpActual = 75
	case "Magicien":
		hpMax = 100
		hpActual = 100
	}

	inventory := []string{} // Inventaire vide au départ

	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		HpMax:     hpMax,
		HpActual:  hpActual,
		Inventory: inventory,
	}
}

func main() {
	character := initCharacter()

	fmt.Println("\n🎮 Personnage créé :")
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nHP: %d/%d\nInventaire: %v\n",
		character.Name, character.Class, character.Level, character.HpActual, character.HpMax, character.Inventory)
}
