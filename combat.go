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
		Name:        "Gobelin d'entraînement",
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
			fmt.Printf("Vous utilisez une %s. %s. PV actuels : %d/%d\n", "Potion de soin", "Vous récupérez 10 PV", p.HP, p.MaxHP)
		},
	}

	return Player{
		Name:      "Héros",
		HP:        30,
		MaxHP:     30,
		Inventory: []Item{potion},
	}
}

func monsterTurn(gobelin *Monster, player *Player, tour int) {
	fmt.Println("\n--- Tour du Monstre ---")

	var degats int
	if tour%3 == 0 {
		degats = gobelin.AttackPower * 2
		fmt.Printf("Attaque SPÉCIALE ! Le %s inflige %d dégâts à %s.\n", gobelin.Name, degats, player.Name)
	} else {
		degats = gobelin.AttackPower
		fmt.Printf("Le %s attaque et inflige %d dégâts à %s.\n", gobelin.Name, degats, player.Name)
	}

	player.HP -= degats
	if player.HP < 0 {
		player.HP = 0
	}

	fmt.Printf("PV restants du joueur : %d/%d\n", player.HP, player.MaxHP)
}

func characterTurn(player *Player, gobelin *Monster, tour int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Tour du Joueur ---")
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choisissez une option : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choix, err := strconv.Atoi(input)

		if err != nil || (choix != 1 && choix != 2) {
			fmt.Println("Choix invalide. Veuillez entrer 1 ou 2.")
			continue
		}

		if choix == 1 {
			degats := 5
			gobelin.CurrentHP -= degats
			if gobelin.CurrentHP < 0 {
				gobelin.CurrentHP = 0
			}

			fmt.Println("\nVous utilisez 'Attaque basique'.")
			fmt.Printf("Vous infligez %d dégâts au %s.\n", degats, gobelin.Name)
			fmt.Printf("PV restants du monstre : %d/%d\n", gobelin.CurrentHP, gobelin.MaxHP)

			if gobelin.CurrentHP > 0 {
				monsterTurn(gobelin, player, tour)
			} else {
				fmt.Printf("\n%s est vaincu !\n", gobelin.Name)
			}

			break

		} else if choix == 2 {
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
				choixItem, err := strconv.Atoi(input)

				if err != nil || choixItem < 0 || choixItem > len(player.Inventory) {
					fmt.Println("Choix invalide.")
					continue
				}

				if choixItem == 0 {
					fmt.Println("Retour au menu.")
					continue
				}

				item := player.Inventory[choixItem-1]
				item.Use(player)

				player.Inventory = append(player.Inventory[:choixItem-1], player.Inventory[choixItem:]...)

				if gobelin.CurrentHP > 0 {
					monsterTurn(gobelin, player, tour)
				}
				break
			}
		}
	}
}

func main() {
	gobelin := initGoblin()
	player := initPlayer()
	tour := 1

	fmt.Println("=== Début du combat ===")
	fmt.Printf("Adversaire : %s (%d PV)\n", gobelin.Name, gobelin.MaxHP)

	for gobelin.CurrentHP > 0 && player.HP > 0 {
		characterTurn(&player, &gobelin, tour)
		tour++
	}

	fmt.Println("\n=== Fin du combat ===")
	if player.HP <= 0 {
		fmt.Println("Vous avez été vaincu...")
	} else {
		fmt.Println("Victoire !")
	}
}
