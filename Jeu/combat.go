package main

import "fmt"


// --- Types ---
type Character2 struct {
	Name      string
	HpActual  int
	HpMax     int
	Attack    int
	Level     int
	CurrentXP int
	MaxXP     int
	Inventory []Item
	Spells    []Spell
}

type Monster struct {
	Name        string
	HpMax       int
	CurrentHP   int
	AttackPower int
	XPReward    int
}

// --- Fonctions d'initialisation ---
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

func initGiant() Monster {
	return Monster{
		Name:        "Géant",
		HpMax:       80,
		CurrentHP:   80,
		AttackPower: 15,
		XPReward:    20,
	}
}

// --- Mécaniques de jeu ---
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

func monsterTurn(monster *Monster, c *Character, turn int) {
	fmt.Println("\n--- Tour du Monstre ---")

	var damage int
	if turn%3 == 0 {
		damage = monster.AttackPower * 2
		fmt.Printf("%s utilise Attaque SPÉCIALE et inflige %d dégâts à %s.\n", monster.Name, damage, c.Name)
	} else {
		damage = monster.AttackPower
		fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", monster.Name, damage, c.Name)
	}

	c.HpActual -= damage
	if c.HpActual < 0 {
		c.HpActual = 0
	}

	fmt.Printf("%s - PV : %d / %d\n", c.Name, c.HpActual, c.HpMax)
}

// --- Combats ---
func trainingFight(c *Character) {
	monster := initGoblin()
	c.Spells = initSpells()
	turn := 1

	fmt.Println("=== Début du Combat d'entraînement ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
	fmt.Printf("Vous : %s - PV : %d / %d | Niveau : %d | XP : %d / %d\n", c.Name, c.HpActual, c.HpMax, c.Level, c.CurrentXP, c.MaxXP)

	for c.HpActual > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			gainXP(c, monster.XPReward)
			break
		}

		monsterTurn(&monster, c, turn)

		if c.HpActual <= 0 {
			fmt.Printf("%s est de retour au lobby !\n", c.Name)
			c.HpActual = c.HpMax / 2
			fmt.Printf("%s est ressuscité avec %d HP.\n", c.Name, c.HpActual)
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat.")
}

func giantFight(c *Character) {
	monster := initGiant()
	c.Spells = initSpells()
	turn := 1

	fmt.Println("=== Combat contre le Géant ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
	fmt.Printf("Vous : %s - PV : %d / %d | Niveau : %d | XP : %d / %d\n", c.Name, c.HpActual, c.HpMax, c.Level, c.CurrentXP, c.MaxXP)

	for c.HpActual > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			gainXP(c, monster.XPReward)
			break
		}

		monsterTurn(&monster, c, turn)

		if c.HpActual <= 0 {
			fmt.Printf("%s est de retour au lobby !\n", c.Name)
			c.HpActual = c.HpMax / 2
			fmt.Printf("%s est ressuscité avec %d HP.\n", c.Name, c.HpActual)
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat.")
}

// --- Affichage du statut ---
func (c *Character) DisplayStatus() {
	fmt.Printf("\n=== Statut de %s ===\n", c.Name)
	fmt.Printf("Niveau : %d\n", c.Level)
	fmt.Printf("PV : %d / %d\n", c.HpActual, c.HpMax)
	fmt.Printf("Attaque : %d\n", c.Attack)
	fmt.Printf("XP : %d / %d\n", c.CurrentXP, c.MaxXP)

	fmt.Println("Sorts connus :")
	for _, s := range c.Spells {
		fmt.Printf("- %s (%d dégâts)\n", s.Name, s.Damage)
	}

	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire : Vide")
	} else {
		fmt.Println("Inventaire :")
		for _, item := range c.Inventory {
			fmt.Printf("- %s : %s\n", item.Name, item.EffectDesc)
		}
	}
	fmt.Println()
}
