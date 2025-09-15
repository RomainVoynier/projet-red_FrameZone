package main

import (
	"fmt"
)

type Personnage struct {
	Nom    string
	Classe string
	Skills []string
}

func apprendreSort(p *Personnage, sort string) {
	for _, skill := range p.Skills {
		if skill == sort {
			fmt.Printf("%s (%s) connaît déjà le sort « %s ».\n", p.Nom, p.Classe, sort)
			return
		}
	}

	p.Skills = append(p.Skills, sort)
	fmt.Printf("✨ %s (%s) a appris le sort « %s » !\n", p.Nom, p.Classe, sort)
}

func attribuerSortParClasse(p *Personnage) {
	switch p.Classe {
	case "Chevalier":
		apprendreSort(p, "Onde de choc")
	case "Archer":
		apprendreSort(p, "Flèche perçante")
	case "Mage":
		apprendreSort(p, "Boule de feu")
	default:
		fmt.Printf("Classe inconnue pour %s\n", p.Nom)
	}
}

func main() {
	chevalier := Personnage{Classe: "Chevalier"}
	archer := Personnage{Classe: "Archer"}
	mage := Personnage{Classe: "Mage"}

	personnages := []*Personnage{&chevalier, &archer, &mage}

	for _, perso := range personnages {
		attribuerSortParClasse(perso)
	}

	fmt.Println("\n Tentative de réapprentissage des sorts :")
	for _, perso := range personnages {
		attribuerSortParClasse(perso)
	}

	fmt.Println("\n Compétences finales :")
	for _, perso := range personnages {
		fmt.Printf("- %s (%s) : %v\n", perso.Nom, perso.Classe, perso.Skills)
	}
}
