package main

import (
	"fmt"
	"strings"
)

// Structure Character
type Character struct {
	Name      string
	Class     string
	Level     int
	HpMax     int
	HpActual  int
	Inventory []string
	Smic      int
}

// Fonction pour forcer un choix de classe valide
func choisirClasse() string {
	for {
		var class string
		fmt.Print("Choisissez une classe (Chevalier, Archer, Magicien) : ")
		fmt.Scanln(&class)

		class = strings.Title(strings.ToLower(class))

		switch class {
		case "Chevalier", "Archer", "Magicien":
			return class
		default:
			fmt.Println("❌ Classe invalide. Veuillez entrer : Chevalier, Archer ou Magicien.")
		}
	}
}

// Fonction qui initialise un personnage
func initCharacter() Character {
	var name string
	fmt.Print("Entrez le nom de votre personnage : ")
	fmt.Scanln(&name)

	class := choisirClasse()
	level := 1
	var hpMax, hpActual, smic int

	switch class {
	case "Chevalier":
		hpMax = 150
		hpActual = 150
		smic = 100
	case "Archer":
		hpMax = 75
		hpActual = 75
		smic = 100
	case "Magicien":
		hpMax = 100
		hpActual = 100
		smic = 100
	}

	inventory := []string{}

	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		HpMax:     hpMax,
		HpActual:  hpActual,
		Inventory: inventory,
		Smic:      smic,
	}
}

// Méthode pour afficher les infos du personnage
func (c Character) displayInfo() {
	fmt.Println("\n📋 Informations du personnage :")
	fmt.Printf("🧾 Nom        : %s\n", c.Name)
	fmt.Printf("🗡️ Classe     : %s\n", c.Class)
	fmt.Printf("📊 Niveau     : %d\n", c.Level)
	fmt.Printf("❤️ HP         : %d/%d\n", c.HpActual, c.HpMax)
	fmt.Printf("🎒 Inventaire : %v\n", c.Inventory)
	fmt.Printf("💰 Smic       : %d\n", c.Smic)
}

// Point d'entrée
func main() {
	character := initCharacter()
	character.displayInfo()
}
