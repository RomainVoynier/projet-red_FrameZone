package main

import (
	"fmt"
)

type PotionDeVie struct {
	Name     string
	Quantite int
}

type MarchandNoir struct {
	Name      string
	Inventory []PotionDeVie
	Usable    bool
}

func accessInventory(p *MarchandNoir) {
	fmt.Println("\n=== Inventaire du Marché Noir ===")

	if len(p.Inventory) == 0 {
		fmt.Println("\n=== Ton inventaire est vide ===")
		return
	}

	for _, potion := range p.Inventory {
		fmt.Printf("Potion : %s | Quantité : %d\n", potion.Name, potion.Quantite)
	}
}

func M() MarchandNoir {
	marchand := MarchandNoir{
		Name:      "Marchand Test",
		Inventory: []PotionDeVie{},
		Usable:    true,
	}

	accessInventory(&marchand)

	marchand.Inventory = append(marchand.Inventory,
		PotionDeVie{"Potion de Vie", 30},
	)

	accessInventory(&marchand)

	return marchand
}
