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

	inventory := []string{} // Inventaire vide au départ

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

// Fonction main obligatoire pour exécuter le programme
func main() {
	character := initCharacter()

	fmt.Println("\n🎮 Personnage créé :")
	fmt.Printf("Nom: %s\nClasse: %s\nNiveau: %d\nHP: %d/%d\nInventaire: %v\nSmic: %d\n",
		character.Name, character.Class, character.Level, character.HpActual, character.HpMax, character.Inventory, character.Smic)
}

// Fonction pour afficher les informations du personnage
func (c Character) displayInfo() {
	fmt.Println("\n📋 Informations du personnage :")
	fmt.Printf("🧾 Nom        : %s\n", c.Name)
	fmt.Printf("🗡️ Classe     : %s\n", c.Class)
	fmt.Printf("📊 Niveau     : %d\n", c.Level)
	fmt.Printf("❤️ HP         : %d/%d\n", c.HpActual, c.HpMax)
	fmt.Printf("🎒 Inventaire : %v\n", c.Inventory)
	fmt.Printf("💰 Smic       : %d pièces\n", c.Smic)
}
