package main

import (
	"bufio"
	"fmt"
	"os"
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

func choisirClasse() Classe {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choisissez une classe :")
		for _, classe := range ClassesDisponibles {
			fmt.Printf("- %s : %s\n", classe.Nom, classe.Description)
		}
		fmt.Print("Votre choix : ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Entrée invalide.")
			continue
		}

		input = strings.TrimSpace(input)
		input = strings.Title(strings.ToLower(input)) // On met la première lettre en majuscule

		if classe, ok := ClassesDisponibles[input]; ok {
			return classe
		}
		fmt.Println("Classe invalide.")
	}
}

func InitCharacter() Character {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom de votre personnage : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	classe := choisirClasse()

	var attackBase int
	var spells []Spell

	// Initialisation selon la classe
	switch classe.Nom {
	case "Chevalier":
		attackBase = 10
		spells = []Spell{
			{Name: "Coup d'estoc", Damage: 20, Used: false},
			{Name: "Passage en Force", Damage: 25, Used: false},
		}
	case "Archer":
		attackBase = 8
		spells = []Spell{
			{Name: "Flèche empoisonnée", Damage: 25, Used: false},
			{Name: "Tir Fatal", Damage: 40, Used: false},
		}
	case "Magicien":
		attackBase = 6
		spells = []Spell{
			{Name: "Blizard Éternel", Damage: 35, Used: false},
			{Name: "Éclair pulverisant", Damage: 25, Used: false},
		}
	}

	return Character{
		Name:       name,
		Class:      classe,
		Level:      1,
		HpMax:      classe.HpMax,
		HpActual:   classe.HpMax,
		Attack:     attackBase,
		Smic:       50,
		Spells:     spells,
		Inventory:  []Item{},
		Equipement: Equipement{},
	}
}


func (c *Character) IsDead() bool {
	if c.HpActual <= 0 {
		fmt.Printf("%s est de retour au lobby ! \n", c.Name)
		c.HpActual = c.HpMax / 2
		fmt.Printf("%s est ressuscité avec %d HP.\n", c.Name, c.HpActual)
		return true
	}
	return false
}

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
