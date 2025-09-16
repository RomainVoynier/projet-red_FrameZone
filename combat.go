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
	MaxHP       int
	CurrentHP   int
	AttackPower int
}

type Item struct {
	Name       string
	EffectDesc string
	Use        func(*Player)
}

type Player struct {
	Name      string
	HP        int
	MaxHP     int
	Inventory []Item
}

func initGoblin() Monster {
	return Monster{
		Name:        "Golem d'entraînement",
		MaxHP:       40,
		CurrentHP:   40,
		AttackPower: 5,
	}
}

func initPlayer() Player {
	potion := Item{
		Name:       "Potion de soin",
		EffectDesc: "Rend 10 PV",
		Use: func(p *Player) {
			heal := 10
			p.HP += heal
			if p.HP > p.MaxHP {
				p.HP = p.MaxHP
			}
			fmt.Printf("Vous utilisez %s. %s. PV actuels : %d/%d\n", "Potion de soin", "Vous récupérez 10 PV", p.HP, p.MaxHP)
		},
	}

	return Player{
		Name:      "Personnage",
		HP:        30,
		MaxHP:     30,
		Inventory: []Item{potion},
	}
}

func goblinPattern(monster *Monster, player *Player, turn int) {
	fmt.Println("\n--- Tour du Monstre ---")

	var damage int
	if turn%3 == 0 {
		damage = monster.AttackPower * 2
		fmt.Printf("%s utilise Attaque SPÉCIALE et inflige %d dégâts à %s.\n", monster.Name, damage, player.Name)
	} else {
		damage = monster.AttackPower
		fmt.Printf("%s attaque et inflige %d dégâts à %s.\n", monster.Name, damage, player.Name)
	}

	player.HP -= damage
	if player.HP < 0 {
		player.HP = 0
	}

	fmt.Printf("%s - PV : %d / %d\n", player.Name, player.HP, player.MaxHP)
}

func charTurn(player *Player, monster *Monster) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Tour du Joueur ---")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choisissez une option : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)

		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Choix invalide. Veuillez entrer 1 ou 2.")
			continue
		}

		if choice == 1 {
			damage := 5
			monster.CurrentHP -= damage
			if monster.CurrentHP < 0 {
				monster.CurrentHP = 0
			}

			fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s.\n", player.Name, damage, monster.Name)
			fmt.Printf("%s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.MaxHP)

			break

		} else if choice == 2 {
			if len(player.Inventory) == 0 {
				fmt.Println("\nInventaire vide.")
			} else {
				fmt.Println("\n--- Inventaire ---")
				for i, item := range player.Inventory {
					fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.EffectDesc)
				}
				fmt.Print("Choisissez un objet à utiliser (ou 0 pour annuler) : ")

				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				itemChoice, err := strconv.Atoi(input)

				if err != nil || itemChoice < 0 || itemChoice > len(player.Inventory) {
					fmt.Println("Choix invalide.")
					continue
				}

				if itemChoice == 0 {
					fmt.Println("Retour au menu.")
					continue
				}

				item := player.Inventory[itemChoice-1]
				fmt.Printf("Vous utilisez %s.\n", item.Name)
				item.Use(player)

				player.Inventory = append(player.Inventory[:itemChoice-1], player.Inventory[itemChoice:]...)

				break
			}
		}
	}
}

func trainingFight() {
	player := initPlayer()
	monster := initGoblin()
	turn := 1

	fmt.Println("=== Début du Combat d'entraînement ===")
	fmt.Printf("Adversaire : %s - PV : %d / %d\n", monster.Name, monster.CurrentHP, monster.MaxHP)
	fmt.Printf("Vous : %s - PV : %d / %d\n", player.Name, player.HP, player.MaxHP)

	for player.HP > 0 && monster.CurrentHP > 0 {
		fmt.Printf("\n=== TOUR %d ===\n", turn)

		charTurn(&player, &monster)

		if monster.CurrentHP <= 0 {
			fmt.Printf("\n%s est vaincu ! Victoire !\n", monster.Name)
			break
		}

		goblinPattern(&monster, &player, turn)

		if player.HP <= 0 {
			fmt.Printf("\n%s a été vaincu ! Défaite...\n", player.Name)
			break
		}

		turn++
	}

	fmt.Println("\nRetour au menu principal...")
}

func mainMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Entraînement")
		fmt.Println("0. Quitter")
		fmt.Print("Choisissez une option : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			trainingFight()
		case "0":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez entrer 1 ou 0.")
		}
	}
}

func main() {
	mainMenu()
}
