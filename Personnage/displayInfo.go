package main

import (
	"fmt"
)
// Fonction pour afficher les informations du personnage
func (c Character) displayInfo() {
	fmt.Println("\n📋 Informations du personnage :")
	fmt.Printf("🧾 Nom        : %s\n", c.Name)
	fmt.Printf("🗡️ Classe     : %s\n", c.Class)
	fmt.Printf("📊 Niveau     : %d\n", c.Level)
	fmt.Printf("❤️ HP         : %d/%d\n", c.HpActual, c.HpMax)
	fmt.Printf("🎒 Inventaire : %v\n", c.Inventory)
	fmt.Printf("💰 Smic       : %d\n", c.Smic)
}

func main() {
	character := initCharacter()
	character.displayInfo()
}
