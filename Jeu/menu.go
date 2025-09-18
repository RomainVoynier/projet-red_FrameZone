package main

import (
	"fmt"
	"time"
)

func main() {
	introduction()
	menu()
}

// === INTRODUCTION DU JEU ===
func introduction() {
	intro := []string{
		"Toi, dernier survivant de cette nuit apocalyptique, as vu ton foyer sombrer sous les flammes, le givre et la foudre.",
		"Depuis ce jour, une seule chose te maintient en vie : la vengeance.",
		"",
		"Dans un monde brisé, rongé par les conflits, les créatures anciennes et les mensonges oubliés, tu poursuis une légende vivante…",
		"Un dragon dont le nom est murmuré avec effroi : Alatreon.",
		"",
		"Ton voyage commence ici.",
		"Les cendres du passé n’ont pas fini de parler.",
	}

	fmt.Println("=== PROLOGUE ===")
	for _, ligne := range intro {
		fmt.Println(ligne)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("\n(Appuyez sur Entrée pour continuer...)")
	var pause string
	fmt.Scanln(&pause)
}

// === FONCTION AFFICHER LES NOMS ===
func afficherNoms() {
	fmt.Println("\n=== Liste des noms ===")
	fmt.Println("1. ABBA")
	fmt.Println("2. Steven Spielberg")
	fmt.Println("\n(Appuyez sur Entrée pour continuer)")
	var pause string
	fmt.Scanln(&pause)
}

// === MENU PRINCIPAL ===
func menu() {
	character := InitCharacter()
	marchand := M()
	hero := convertCharacterToHero(&character)
	trainingFight(hero)

	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les infos du personnage")
		fmt.Println("2. Aller chez le Forgeron")
		fmt.Println("3. Qui sont-ils")
		fmt.Println("4. Entrainement")
		fmt.Println("5. Marchand")
		fmt.Println("6. Quitter")

		var choix int
		fmt.Print("Choix : ")
		_, err := fmt.Scanln(&choix)
		if err != nil {
			fmt.Println("Entrée invalide.")
			continue
		}

		switch choix {
		case 1:
			character.DisplayInfo()
		case 2:
			ForgeronMenu(&character)
		case 3:
			afficherNoms()
		case 4:
			trainingFight(hero)
		case 5:
			accessInventory(&marchand)
		case 6:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func convertCharacterToHero(c *Character) *Hero {
	return &Hero{
		Name:     c.Name,
		HpActual: c.HpActual,
		HpMax:    c.HpMax,
		Level:    c.Level,
		// Complète avec les autres champs nécessaires
	}
}
