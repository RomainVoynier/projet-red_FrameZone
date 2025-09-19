package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- Types nécessaires (à garder dans le même package) ---

type Monster struct {
	Name        string
	HpMax       int
	CurrentHP   int
	AttackPower int
	XPReward    int
}

// --- Initialisation des monstres ---

func initGoblin() Monster {
	return Monster{
		Name:        "Alatreon",
		HpMax:       150,
		CurrentHP:   150,
		AttackPower: 20,
		XPReward:    30,
	}
}

func initGiant() Monster {
	return Monster{
		Name:        "Géant",
		HpMax:       80,
		CurrentHP:   80,
		AttackPower: 15,
		XPReward:    20,
	}
}

func initDragon() Monster {
	return Monster{
		Name:        "Dragon",
		HpMax:       150,
		CurrentHP:   150,
		AttackPower: 30,
		XPReward:    40,
	}
}

// --- Mécanique d'expérience ---

func gainXP(c *Character, amount int) {
	fmt.Printf("\nVous gagnez %d points d'expérience !\n", amount)
	c.CurrentXP += amount

	for c.CurrentXP >= c.MaxXP {
		c.CurrentXP -= c.MaxXP
		c.Level++
		c.MaxXP += 10
		c.HpMax += 5
		c.Attack += 2
		c.HpActual = c.HpMax

		fmt.Printf("\n🎉 Vous passez au niveau %d !\n", c.Level)
		fmt.Printf("→ PV max : %d | Attaque : %d | XP pour le prochain niveau : %d\n", c.HpMax, c.Attack, c.MaxXP)
	}

	fmt.Printf("XP actuelle : %d / %d\n", c.CurrentXP, c.MaxXP)
}

// --- Tour du monstre ---

func monsterTurn(m *Monster, c *Character, turn int) {
	fmt.Println("\n--- Tour du Monstre ---")

	var damage int
	if turn%3 == 0 {
		damage = m.AttackPower * 2
		fmt.Printf("%s utilise Attaque SPÉCIALE et inflige %d dégâts à %s.\n", m.Name, damage, c.Name)
	} else {
		damage = m.AttackPower
		fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", m.Name, damage, c.Name)
	}

	c.HpActual -= damage
	if c.HpActual < 0 {
		c.HpActual = 0
	}

	fmt.Printf("%s - PV : %d / %d\n", c.Name, c.HpActual, c.HpMax)
}

// --- Tour du joueur ---

func heroTurn(c *Character, m *Monster) {
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
			damage := c.Attack
			m.CurrentHP -= damage
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}
			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s.\n", c.Name, damage, m.Name)
			fmt.Printf("%s - PV : %d / %d\n", m.Name, m.CurrentHP, m.HpMax)
			return

		case 2:
			availableSpells := []int{}
			fmt.Println("\n--- Sorts disponibles ---")
			for i, spell := range c.Spells {
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
			spell := &c.Spells[spellIndex]

			m.CurrentHP -= spell.Damage
			if m.CurrentHP < 0 {
				m.CurrentHP = 0
			}

			fmt.Printf("%s lance %s et inflige %d dégâts à %s.\n", c.Name, spell.Name, spell.Damage, m.Name)
			fmt.Printf("%s - PV : %d / %d\n", m.Name, m.CurrentHP, m.HpMax)

			spell.Used = true
			return

		case 3:
			if len(c.Inventory) == 0 {
				fmt.Println("\nInventaire vide.")
			} else {
				fmt.Println("\n--- Inventaire ---")
				for i, item := range c.Inventory {
					fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.EffectDesc)
				}
				fmt.Print("Choisissez un objet à utiliser (ou 0 pour annuler) : ")

				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				itemChoice, err := strconv.Atoi(input)

				if err != nil || itemChoice < 0 || itemChoice > len(c.Inventory) {
					fmt.Println("Choix invalide.")
					continue
				}

				if itemChoice == 0 {
					fmt.Println("Retour au menu.")
					continue
				}

				item := c.Inventory[itemChoice-1]
				fmt.Printf("Vous utilisez %s.\n", item.Name)
				item.Use(c)

				// Retirer l'objet de l'inventaire
				c.Inventory = append(c.Inventory[:itemChoice-1], c.Inventory[itemChoice:]...)
				return
			}
		}
	}
}

// --- Combat d'entraînement ---

func trainingFight(c *Character) {
	monster := initGoblin()
	// Initialiser les sorts à chaque combat
	for i := range c.Spells {
		c.Spells[i].Used = false
	}
	turn := 1

	fmt.Println("=== Début du combat contre Alatreon ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
	fmt.Printf("Vous : %s - PV : %d / %d | Niveau : %d | XP : %d / %d\n", c.Name, c.HpActual, c.HpMax, c.Level, c.CurrentXP, c.MaxXP)

	for c.HpActual > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)

		heroTurn(c, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			gainXP(c, monster.XPReward)
			break
		}

		monsterTurn(&monster, c, turn)

		if c.HpActual <= 0 {
			c.HpActual = c.HpMax / 2
			fmt.Printf("%s est ressuscité avec %d HP.\n", c.Name, c.HpActual)
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat.")
}
