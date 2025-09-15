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
			fmt.Println("Classe invalide. Veuillez entrer : Chevalier, Archer ou Magicien.")
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
	case "Archer":
		hpMax = 75
	case "Magicien":
		hpMax = 100
	}
	hpActual = hpMax
	smic = 50 // Valeur de départ pour les tests

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
	fmt.Println("\nInformations du personnage :")
	fmt.Printf(" Nom        : %s\n", c.Name)
	fmt.Printf(" Classe     : %s\n", c.Class)
	fmt.Printf(" Niveau     : %d\n", c.Level)
	fmt.Printf(" HP         : %d/%d\n", c.HpActual, c.HpMax)
	fmt.Printf(" Inventaire : %v\n", c.Inventory)
	fmt.Printf(" Smic       : %d\n", c.Smic)
}

// Méthode pour accéder à l'inventaire du personnage
func (c Character) accessInventory() {
	fmt.Println("\nInventaire du personnage :")
	if len(c.Inventory) == 0 {
		fmt.Println("L'inventaire est vide.")
		return
	}
	for i, item := range c.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

// Fonction du forgeron
func forgeronMenu(c *Character) {
	for {
		fmt.Println("\nBienvenue chez le Forgeron")
		fmt.Println("1. Couronne de Lauriers : 5 Smic")
		fmt.Println("2. Tronc d'Arbre : 15 Smic")
		fmt.Println("3. Bottes de Sapin : 10 Smic")
		fmt.Println("4. Retour")

		var choix int
		fmt.Print("Choix : ")
		if _, err := fmt.Scanln(&choix); err != nil {
			fmt.Println("Entrée invalide.")
			continue
		}

		var item string
		var cost int

		switch choix {
		case 1:
			item = "Couronne de Lauriers"
			cost = 5
		case 2:
			item = "Tronc d'Arbre"
			cost = 15
		case 3:
			item = "Bottes de Sapin"
			cost = 10
		case 4:
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if c.Smic >= cost {
			c.Smic -= cost
			c.Inventory = append(c.Inventory, item)
			fmt.Printf("%s fabriqué et ajouté à votre inventaire.\n", item)
			fmt.Printf("Smic restant : %d\n", c.Smic)
		} else {
			fmt.Println("Pas assez de smic pour fabriquer cet objet.")
		}
	}
}


// Point d'entrée
func main() {
	character := initCharacter()

	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les infos du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Aller chez le Forgeron")
		fmt.Println("4. Quitter")

		var choix int
		fmt.Print("Choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			character.displayInfo()
		case 2:
			character.accessInventory()
		case 3:
			forgeronMenu(&character)
		case 4:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
