package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monster struct {
	Name        string
	HpMax       int
	CurrentHP   int
	AttackPower int
	XPReward    int
}

// Réinitialise les sorts à usage unique
func initSpells() []Spell {
	return []Spell{
		{Name: "Coup de poing", Damage: 8, Used: false},
		{Name: "Boule de feu", Damage: 18, Used: false},
	}
}

func initGoblin() Monster {
	return Monster{
		Name:        "Golem d'entraînement",
		HpMax:       40,
		CurrentHP:   40,
		AttackPower: 5,
		XPReward:    12,
	}
}

func initPlayer() *Character {
	potion := Item{
		Name:       "Potion de soin",
		EffectDesc: "Rend 10 PV",
		Use: func(c *Character) {
			heal := 10
			c.HpActual += heal
			if c.HpActual > c.HpMax {
				c.HpActual = c.HpMax
			}
			fmt.Printf("Vous utilisez %s. Vous récupérez 10 PV. PV actuels : %d/%d\n", "Potion de soin", c.HpActual, c.HpMax)
		},
	}

	return &Character{
		Name:      "Personnage",
		HpActual:  30,
		HpMax:     30,
		Inventory: []Item{potion},
		Level:     1,
		CurrentXP: 0,
		MaxXP:     20,
		Attack:    5,
		Spells:    initSpells(),
	}
}

func gainXP(character *Character, amount int) {
	fmt.Printf("\nVous gagnez %d points d'expérience !\n", amount)
	character.CurrentXP += amount

	for character.CurrentXP >= character.MaxXP {
		character.CurrentXP -= character.MaxXP
		character.Level++
		character.MaxXP += 10

		character.HpMax += 5
		character.Attack += 2
		character.HpActual = character.HpMax

		fmt.Printf("\n Vous passez au niveau %d !\n", character.Level)
		fmt.Printf("→ PV max : %d | Attaque : %d | XP pour le prochain niveau : %d\n", character.HpMax, character.Attack, character.MaxXP)
	}

	fmt.Printf("XP actuelle : %d / %d\n", character.CurrentXP, character.MaxXP)
}

func goblinPattern(monster *Monster, character *Character, turn int) {
	fmt.Println("\n--- Tour du Monstre ---")

	var damage int
	if turn%3 == 0 {
		damage = monster.AttackPower * 2
		fmt.Printf("%s utilise Attaque SPÉCIALE et inflige %d dégâts à %s.\n", monster.Name, damage, character.Name)
	} else {
		damage = monster.AttackPower
		fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", monster.Name, damage, character.Name)
	}

	character.HpActual -= damage
	if character.HpActual < 0 {
		character.HpActual = 0
	}

	fmt.Printf("%s - PV : %d / %d\n", character.Name, character.HpActual, character.HpMax)
}

func charTurn(character *Character, monster *Monster) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Tour du Joueur ---")
		fmt.Println("1. Attaquer (attaque basique)")
		fmt.Println("2. Sorts")
		fmt.Println("3. Inventaire")
		fmt.Print("Choisissez une option : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)

		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Choix invalide. Veuillez entrer 1, 2 ou 3.")
			continue
		}

		switch choice {
		case 1:
			damage := character.Attack
			monster.CurrentHP -= damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}

			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s.\n", character.Name, damage, monster.Name)
			fmt.Printf("%s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
			return

		case 2:
			availableSpells := []int{}
			fmt.Println("\n--- Sorts disponibles ---")
			for i, spell := range character.Spells {
				if !spell.Used {
					fmt.Printf("%d. %s (%d dégâts)\n", len(availableSpells)+1, spell.Name, spell.Damage)
					availableSpells = append(availableSpells, i)
				}
			}

			if len(availableSpells) == 0 {
				fmt.Println("Aucun sort disponible.")
				continue
			}

			fmt.Print("Choisissez un sort : ")
			spellInput, _ := reader.ReadString('\n')
			spellInput = strings.TrimSpace(spellInput)
			spellChoice, err := strconv.Atoi(spellInput)

			if err != nil || spellChoice < 1 || spellChoice > len(availableSpells) {
				fmt.Println("Sort invalide.")
				continue
			}

			spellIndex := availableSpells[spellChoice-1]
			spell := &character.Spells[spellIndex]

			monster.CurrentHP -= spell.Damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}

			fmt.Printf("%s lance %s et inflige %d dégâts à %s.\n", character.Name, spell.Name, spell.Damage, monster.Name)
			fmt.Printf("%s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)

			spell.Used = true
			return

		case 3:
			if len(character.Inventory) == 0 {
				fmt.Println("\nInventaire vide.")
			} else {
				fmt.Println("\n--- Inventaire ---")
				for i, item := range character.Inventory {
					fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.EffectDesc)
				}
				fmt.Print("Choisissez un objet à utiliser (ou 0 pour annuler) : ")

				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				itemChoice, err := strconv.Atoi(input)

				if err != nil || itemChoice < 0 || itemChoice > len(character.Inventory) {
					fmt.Println("Choix invalide.")
					continue
				}

				if itemChoice == 0 {
					fmt.Println("Retour au menu.")
					continue
				}

				item := character.Inventory[itemChoice-1]
				fmt.Printf("Vous utilisez %s.\n", item.Name)
				item.Use(character)

				character.Inventory = append(character.Inventory[:itemChoice-1], character.Inventory[itemChoice:]...)
				return
			}
		}
	}
}

func trainingFight(character *Character) {
	monster := initGoblin()
	character.Spells = initSpells() // Réinitialise les sorts à chaque combat
	turn := 1

	fmt.Println("=== Début du Combat d'entraînement ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
	fmt.Printf("Vous : %s - PV : %d / %d | Niveau : %d | XP : %d / %d\n", character.Name, character.HpActual, character.HpMax, character.Level, character.CurrentXP, character.MaxXP)

	for character.HpActual > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)

		charTurn(character, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			gainXP(character, monster.XPReward)
			break
		}

		goblinPattern(&monster, character, turn)

		if character.HpActual <= 0 {
			fmt.Printf("%s est de retour au lobby ! !\n", character.Name)
			character.HpActual = character.HpMax / 2
			fmt.Printf("%s est ressuscité avec %d HP.\n", character.Name, character.HpActual)
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat. Merci d'avoir joué !")
}
