package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- Types ---
type Hero struct {
	Name      string
	HpActual  int
	HpMax     int
	Attack    int
	Level     int
	CurrentXP int
	MaxXP     int
	Inventory []Artifact
	Spells    []Skill
}

type Skill struct {
	Name   string
	Damage int
	Used   bool
}

type Artifact struct {
	Name       string
	EffectDesc string
	Use        func(*Hero)
}

type Monster struct {
	Name        string
	HpMax       int
	CurrentHP   int
	AttackPower int
	XPReward    int
}

// --- Fonctions d'initialisation ---
func initSpells() []Skill {
	return []Skill{
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

func initDragon() Monster {
	return Monster{
		Name:        "Dragon",
		HpMax:       150,
		CurrentHP:   150,
		AttackPower: 30,
		XPReward:    40,
	}
}

func InitHero() *Hero {
	potion := Artifact{
		Name:       "Potion de soin",
		EffectDesc: "Rend 10 PV",
		Use: func(h *Hero) {
			heal := 10
			h.HpActual += heal
			if h.HpActual > h.HpMax {
				h.HpActual = h.HpMax
			}
			fmt.Printf("Vous utilisez %s. Vous récupérez 10 PV. PV actuels : %d/%d\n", "Potion de soin", h.HpActual, h.HpMax)
		},
	}

	return &Hero{
		Name:      "Personnage",
		HpActual:  30,
		HpMax:     30,
		Inventory: []Artifact{potion},
		Level:     1,
		CurrentXP: 0,
		MaxXP:     20,
		Attack:    5,
		Spells:    initSpells(),
	}
}

// --- Mécaniques de jeu ---
func gainXP(hero *Hero, amount int) {
	fmt.Printf("\nVous gagnez %d points d'expérience !\n", amount)
	hero.CurrentXP += amount

	for hero.CurrentXP >= hero.MaxXP {
		hero.CurrentXP -= hero.MaxXP
		hero.Level++
		hero.MaxXP += 10
		hero.HpMax += 5
		hero.Attack += 2
		hero.HpActual = hero.HpMax

		fmt.Printf("\n🎉 Vous passez au niveau %d !\n", hero.Level)
		fmt.Printf("→ PV max : %d | Attaque : %d | XP pour le prochain niveau : %d\n", hero.HpMax, hero.Attack, hero.MaxXP)
	}

	fmt.Printf("XP actuelle : %d / %d\n", hero.CurrentXP, hero.MaxXP)
}

func monsterTurn(monster *Monster, hero *Hero, turn int) {
	fmt.Println("\n--- Tour du Monstre ---")

	var damage int
	if turn%3 == 0 {
		damage = monster.AttackPower * 2
		fmt.Printf("%s utilise Attaque SPÉCIALE et inflige %d dégâts à %s.\n", monster.Name, damage, hero.Name)
	} else {
		damage = monster.AttackPower
		fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", monster.Name, damage, hero.Name)
	}

	hero.HpActual -= damage
	if hero.HpActual < 0 {
		hero.HpActual = 0
	}

	fmt.Printf("%s - PV : %d / %d\n", hero.Name, hero.HpActual, hero.HpMax)
}

func heroTurn(hero *Hero, monster *Monster) {
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
			damage := hero.Attack
			monster.CurrentHP -= damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}
			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s.\n", hero.Name, damage, monster.Name)
			fmt.Printf("%s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
			return

		case 2:
			availableSpells := []int{}
			fmt.Println("\n--- Sorts disponibles ---")
			for i, spell := range hero.Spells {
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
			spell := &hero.Spells[spellIndex]

			monster.CurrentHP -= spell.Damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}

			fmt.Printf("%s lance %s et inflige %d dégâts à %s.\n", hero.Name, spell.Name, spell.Damage, monster.Name)
			fmt.Printf("%s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)

			spell.Used = true
			return

		case 3:
			if len(hero.Inventory) == 0 {
				fmt.Println("\nInventaire vide.")
			} else {
				fmt.Println("\n--- Inventaire ---")
				for i, item := range hero.Inventory {
					fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.EffectDesc)
				}
				fmt.Print("Choisissez un objet à utiliser (ou 0 pour annuler) : ")

				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				itemChoice, err := strconv.Atoi(input)

				if err != nil || itemChoice < 0 || itemChoice > len(hero.Inventory) {
					fmt.Println("Choix invalide.")
					continue
				}

				if itemChoice == 0 {
					fmt.Println("Retour au menu.")
					continue
				}

				item := hero.Inventory[itemChoice-1]
				fmt.Printf("Vous utilisez %s.\n", item.Name)
				item.Use(hero)

				hero.Inventory = append(hero.Inventory[:itemChoice-1], hero.Inventory[itemChoice:]...)
				return
			}
		}
	}
}

// --- Combats ---
func trainingFight(hero *Hero) {
	monster := initGoblin()
	hero.Spells = initSpells()
	turn := 1

	fmt.Println("=== Début du Combat d'entraînement ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
	fmt.Printf("Vous : %s - PV : %d / %d | Niveau : %d | XP : %d / %d\n", hero.Name, hero.HpActual, hero.HpMax, hero.Level, hero.CurrentXP, hero.MaxXP)

	for hero.HpActual > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)

		heroTurn(hero, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			gainXP(hero, monster.XPReward)
			break
		}

		monsterTurn(&monster, hero, turn)

		if hero.HpActual <= 0 {
			fmt.Printf("%s est de retour au lobby !\n", hero.Name)
			hero.HpActual = hero.HpMax / 2
			fmt.Printf("%s est ressuscité avec %d HP.\n", hero.Name, hero.HpActual)
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat.")
}

func giantFight(hero *Hero) {
	monster := initGiant()
	hero.Spells = initSpells()
	turn := 1

	fmt.Println("=== Combat contre le Géant ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
	fmt.Printf("Vous : %s - PV : %d / %d | Niveau : %d | XP : %d / %d\n", hero.Name, hero.HpActual, hero.HpMax, hero.Level, hero.CurrentXP, hero.MaxXP)

	for hero.HpActual > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)
		heroTurn(hero, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			gainXP(hero, monster.XPReward)
			break
		}

		monsterTurn(&monster, hero, turn)

		if hero.HpActual <= 0 {
			fmt.Printf("%s est de retour au lobby !\n", hero.Name)
			hero.HpActual = hero.HpMax / 2
			fmt.Printf("%s est ressuscité avec %d HP.\n", hero.Name, hero.HpActual)
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat.")
}

func dragonFight(hero *Hero) {
	monster := initDragon()
	hero.Spells = initSpells()
	turn := 1

	fmt.Println("=== Combat contre le Dragon ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.HpMax)
	fmt.Printf("Vous : %s - PV : %d / %d | Niveau : %d | XP : %d / %d\n", hero.Name, hero.HpActual, hero.HpMax, hero.Level, hero.CurrentXP, hero.MaxXP)

	for hero.HpActual > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)
		heroTurn(hero, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			gainXP(hero, monster.XPReward)
			break
		}

		monsterTurn(&monster, hero, turn)

		if hero.HpActual <= 0 {
			fmt.Printf("%s est de retour au lobby !\n", hero.Name)
			hero.HpActual = hero.HpMax / 2
			fmt.Printf("%s est ressuscité avec %d HP.\n", hero.Name, hero.HpActual)
			break
		}

		turn++
	}

	fmt.Println("\nFin du combat.")
}

// --- Menu principal ---
func gameMenu() {
	reader := bufio.NewReader(os.Stdin)
	hero := InitHero()

	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Combat d'entraînement (Golem)")
		fmt.Println("2. Combat intermédiaire (Géant)")
		fmt.Println("3. Combat final (Dragon)")
		fmt.Println("4. Quitter")
		fmt.Print("Choisissez une option : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			trainingFight(hero)
		case "2":
			giantFight(hero)
		case "3":
			dragonFight(hero)
		case "4":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func MAIN() {
	gameMenu()
}
