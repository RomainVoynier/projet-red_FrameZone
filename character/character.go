package main

import (
	"fmt"
	"strings"
)

// Structure représentant une classe de personnage
type Classe struct {
	Nom         string
	Description string
	HpMax       int
}

// Map des classes disponibles
var ClassesDisponibles = map[string]Classe{
	"Chevalier": {
		Nom:         "Chevalier",
		Description: "Grand,fort,puissant,résistant mais lent. C'est déja pas mal. ",
		HpMax:       150,
	},
	"Archer": {
		Nom:         "Archer",
		Description: "Expert en flèches… parfois dans le vide",
		HpMax:       75,
	},
	"Magicien": {
		Nom:         "Magicien",
		Description: "Qui n'a jamais révélé d'être Harry Potter ?",
		HpMax:       100,
	},
}

// Structure Character sans inventaire
type Character struct {
	Name     string
	Class    Classe
	Level    int
	HpMax    int
	HpActual int
	Smic     int
}

// Choix de classe avec affichage des descriptions
func choisirClasse() Classe {
	for {
		fmt.Println("Choisissez une classe :")
		for _, classe := range ClassesDisponibles {
			fmt.Printf("- %s : %s\n", classe.Nom, classe.Description)
		}
		fmt.Print("Votre choix : ")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Entrée invalide, veuillez réessayer.")
			continue
		}

		input = strings.Title(strings.ToLower(strings.TrimSpace(input)))

		if classe, ok := ClassesDisponibles[input]; ok {
			return classe
		}
		fmt.Println("Classe invalide. Veuillez entrer : Chevalier, Archer ou Magicien.")
	}
}

// Initialisation personnage sans inventaire
func initCharacter() Character {
	fmt.Print("Entrez le nom de votre personnage : ")
	var name string
	fmt.Scanln(&name)

	classe := choisirClasse()

	level := 1
	hpMax := classe.HpMax
	hpActual := hpMax
	smic := 50

	return Character{
		Name:     name,
		Class:    classe,
		Level:    level,
		HpMax:    hpMax,
		HpActual: hpActual,
		Smic:     smic,
	}
}

// Affichage infos du personnage (sans inventaire)
func (c Character) displayInfo() {
	fmt.Println("\nInformations du personnage :")
	fmt.Printf(" Nom        : %s\n", c.Name)
	fmt.Printf(" Classe     : %s\n", c.Class.Nom)
	fmt.Printf(" Description: %s\n", c.Class.Description)
	fmt.Printf(" Niveau     : %d\n", c.Level)
	fmt.Printf(" HP         : %d/%d\n", c.HpActual, c.HpMax)
	fmt.Printf(" Smic       : %d\n", c.Smic)
}

// Forgeron simplifié sans inventaire
func forgeronMenu(c *Character) {
	for {
		fmt.Println("\nBienvenue chez le Forgeron")
		fmt.Println("1. Couronne de Lauriers : 5 Smic")
		fmt.Println("2. Tronc d'Arbre : 15 Smic")
		fmt.Println("3. Bottes de Sapin : 10 Smic")
		fmt.Println("4. Retour")

		var choix int
		fmt.Print("Choix : ")
		_, err := fmt.Scanln(&choix)
		if err != nil {
			fmt.Println("Entrée invalide.")
			continue
		}

		var cost int
		switch choix {
		case 1:
			cost = 5
		case 2:
			cost = 15
		case 3:
			cost = 10
		case 4:
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if c.Smic >= cost {
			c.Smic -= cost
			fmt.Printf("Objet fabriqué ! Smic restant : %d\n", c.Smic)
		} else {
			fmt.Println("Trop pauvre pour cet objet.")
		}
	}
}

// Afficher 2 noms sans interaction
func afficherNoms() {
	fmt.Println("\n=== Liste des noms ===")
	fmt.Println("1.ABBA")
	fmt.Println("2. Steven Spielberg")
	fmt.Println("\n(Appuyez sur Entrée pour continuer)")
	fmt.Scanln()
}

func main() {
	character := initCharacter()

	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les infos du personnage")
		fmt.Println("2. Aller chez le Forgeron")
		fmt.Println("3. Qui sont-ils ")
		fmt.Println("4. Quitter")

		var choix int
		fmt.Print("Choix : ")
		_, err := fmt.Scanln(&choix)
		if err != nil {
			fmt.Println("Entrée invalide.")
			continue
		}

		switch choix {
		case 1:
			character.displayInfo()
		case 2:
			forgeronMenu(&character)
		case 3:
			afficherNoms()
		case 4:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
