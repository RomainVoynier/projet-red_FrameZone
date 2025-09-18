package main

import "fmt"

var Boutique = []EquipementPiece{
	{Nom: "Couronne de Lauriers", Cout: 5, Slot: "Tete", BonusHP: 5},
	{Nom: "Tronc d'Arbre", Cout: 15, Slot: "Torse", BonusHP: 20},
	{Nom: "Bottes de Sapin", Cout: 10, Slot: "Pieds", BonusHP: 10},
}

// Menu du forgeron : Achat d'équipement
func ForgeronMenu(c *Character) {
	for {
		fmt.Println("\nBienvenue chez le Forgeron")
		fmt.Printf("Smic actuel : %d\n", c.Smic)
		fmt.Println("Objets disponibles :")
		for i, obj := range Boutique {
			fmt.Printf("%d. %s (%s) : %d Smic, +%d HP\n", i+1, obj.Nom, obj.Slot, obj.Cout, obj.BonusHP)
		}
		fmt.Println("0. Retour")

		var choix int
		fmt.Print("Choix : ")
		_, err := fmt.Scanln(&choix)
		if err != nil || choix < 0 || choix > len(Boutique) {
			fmt.Println("Choix invalide.")
			continue
		}

		if choix == 0 {
			fmt.Println("Retour au menu principal.")
			return
		}

		objet := Boutique[choix-1]
		if c.Smic < objet.Cout {
			fmt.Println("Trop pauvre pour cet objet.")
			continue
		}

		// Achat
		c.Smic -= objet.Cout

		// Équipement selon le slot
		switch objet.Slot {
		case "Tete":
			c.Equipement.Tete = &objet
		case "Torse":
			c.Equipement.Torse = &objet
		case "Pieds":
			c.Equipement.Pieds = &objet
		default:
			fmt.Println("Slot inconnu, impossible d’équiper.")
			continue
		}

		// Mise à jour des HP max et actuels
		c.HpMax = c.CalculerHpMax()
		if c.HpActual > c.HpMax {
			c.HpActual = c.HpMax
		}

		fmt.Printf("Vous avez équipé %s ! Nouveau HP max : %d\n", objet.Nom, c.HpMax)
	}
}
