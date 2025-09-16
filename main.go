package main

import (
	"fmt"
	"framezone/character"
	"framezone/potion"
	"framezone/menu"
)

func main() {
	fmt.Println("Bienvenue dans FrameZone !")

	player := character.CharacterCreation()
	character.DisplayCharacterInfo(player)

	menu.AfficherMenu(&player)

	fmt.Println("Fin du jeu.")
}
