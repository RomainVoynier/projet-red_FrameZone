package main

import "fmt"

// Structure Equipement
type Equipement struct {
	Tete  string
	Torse string
	Pieds string
}

// Menu du forgeron : Achat d'équipement
func ForgeronMenu(c *Character) {
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
			if c.Smic >= cost {
				c.Smic -= cost
				c.Equipement.Tete = "Couronne de Lauriers"
				fmt.Printf("Objet équipé sur la tête ! Smic restant : %d\n", c.Smic)
			} else {
				fmt.Println("Trop pauvre pour cet objet.")
			}
		case 2:
			cost = 15
			if c.Smic >= cost {
				c.Smic -= cost
				c.Equipement.Torse = "Tronc d'Arbre"
				fmt.Printf("Objet équipé sur le torse ! Smic restant : %d\n", c.Smic)
			} else {
				fmt.Println("Trop pauvre pour cet objet.")
			}
		case 3:
			cost = 10
			if c.Smic >= cost {
				c.Smic -= cost
				c.Equipement.Pieds = "Bottes de Sapin"
				fmt.Printf("Objet équipé sur les pieds ! Smic restant : %d\n", c.Smic)
			} else {
				fmt.Println("Trop pauvre pour cet objet.")
			}
		case 4:
			fmt.Println("Retour au menu principal.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
