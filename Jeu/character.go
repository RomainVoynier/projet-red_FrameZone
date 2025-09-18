package main

import (
	"fmt"
	"strings"
)

type Item struct {
	Name       string
	EffectDesc string
	Use        func(*Character)
}

type Spell struct {
	Name   string
	Damage int
	Used   bool
}

type EquipementPiece struct {
	Nom     string
	Cout    int
	Slot    string // Tete, Torse, Pieds
	BonusHP int
}

type Equipement struct {
	Tete  *EquipementPiece
	Torse *EquipementPiece
	Pieds *EquipementPiece
}

type Classe struct {
	Nom         string
	Description string
	HpMax       int
}

type Character struct {
	Name       string
	HpMax      int
	HpActual   int
	Level      int
	CurrentXP  int
	MaxXP      int
	Attack     int
	Inventory  []Item
	Spells     []Spell
	Class      Classe
	Smic       int
	Equipement Equipement
}

// Map des classes disponibles
var ClassesDisponibles = map[string]Classe{
	"Chevalier": {
		Nom:         "Chevalier",
		Description: "Grand, fort, puissant, résistant mais lent. C'est déja pas mal.",
		HpMax:       150,
	},
	"Archer": {
		Nom:         "Archer",
		Description: "Expert en flèches… parfois dans le vide.",
		HpMax:       75,
	},
	"Magicien": {
		Nom:         "Magicien",
		Description: "Qui n'a jamais rêvé d'être Harry Potter ?",
		HpMax:       100,
	},
}

// Fonction pour choisir une classe
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
			fmt.Println("Entrée invalide.")
			continue
		}

		input = strings.Title(strings.ToLower(strings.TrimSpace(input)))

		if classe, ok := ClassesDisponibles[input]; ok {
			return classe
		}
		fmt.Println("Classe invalide.")
	}
}

// Initialisation du personnage
func InitCharacter() Character {
	fmt.Print("Entrez le nom de votre personnage : ")
	var name string
	fmt.Scanln(&name)

	classe := choisirClasse()

	var attackBase int
	switch classe.Nom {
	case "Chevalier":
		attackBase = 10
	case "Archer":
		attackBase = 8
	case "Magicien":
		attackBase = 6
		// Pas de default
	}

	return Character{
		Name:       name,
		Class:      classe,
		Level:      1,
		HpMax:      classe.HpMax,
		HpActual:   classe.HpMax,
		Smic:       50,
		Equipement: Equipement{},
		Attack:     attackBase,
	}
}

// Vérifie si le personnage est mort et le ressuscite
func (c *Character) IsDead() bool {
	if c.HpActual <= 0 {
		fmt.Printf("%s est de retour au lobby ! !\n", c.Name)
		c.HpActual = c.HpMax / 2
		fmt.Printf("%s est ressuscité avec %d HP.\n", c.Name, c.HpActual)
		return true
	}
	return false
}

func (c *Character) ApprendreSort(sort string, damage int) {
	for _, s := range c.Spells {
		if s.Name == sort {
			fmt.Printf("%s connaît déjà le sort %s.\n", c.Name, sort)
			return
		}
	}
	c.Spells = append(c.Spells, Spell{Name: sort, Damage: damage})
	fmt.Printf("%s a appris le sort %s avec %d de dégâts.\n", c.Name, sort, damage)
}

// Affichage des infos du personnage
func (c Character) DisplayInfo() {
	fmt.Println("\nInformations du personnage :")
	fmt.Printf(" Nom        : %s\n", c.Name)
	fmt.Printf(" Classe     : %s\n", c.Class.Nom)
	fmt.Printf(" Niveau     : %d\n", c.Level)
	fmt.Printf(" HP         : %d/%d\n", c.HpActual, c.HpMax)
	fmt.Printf(" Smic       : %d\n", c.Smic)
	fmt.Println(" Équipement :")

	if c.Equipement.Tete != nil {
		fmt.Printf("  Tête  : %s (+%d HP)\n", c.Equipement.Tete.Nom, c.Equipement.Tete.BonusHP)
	} else {
		fmt.Println("  Tête  : [Aucun]")
	}

	if c.Equipement.Torse != nil {
		fmt.Printf("  Torse : %s (+%d HP)\n", c.Equipement.Torse.Nom, c.Equipement.Torse.BonusHP)
	} else {
		fmt.Println("  Torse : [Aucun]")
	}

	if c.Equipement.Pieds != nil {
		fmt.Printf("  Pieds : %s (+%d HP)\n", c.Equipement.Pieds.Nom, c.Equipement.Pieds.BonusHP)
	} else {
		fmt.Println("  Pieds : [Aucun]")
	}
}

func (c *Character) CalculerHpMax() int {
	baseHP := c.Class.HpMax
	bonus := 0

	if c.Equipement.Tete != nil {
		bonus += c.Equipement.Tete.BonusHP
	}
	if c.Equipement.Torse != nil {
		bonus += c.Equipement.Torse.BonusHP
	}
	if c.Equipement.Pieds != nil {
		bonus += c.Equipement.Pieds.BonusHP
	}

	return baseHP + bonus
}
