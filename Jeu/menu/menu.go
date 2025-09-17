package menu

import (
	"fmt"
    "projet-red_FrameZone/Jeu/character"
)

// Affichage simple de noms
func afficherNoms() {
	fmt.Println("\n=== Liste des noms ===")
	fmt.Println("1. ABBA")
	fmt.Println("2. Steven Spielberg")
	fmt.Println("\n(Appuyez sur Entrée pour continuer)")
	var pause string
	fmt.Scanln(&pause)
}


// Simulation d'entraînement

func training() {
	fmt.Println("\n>>> Début de l'entraînement <<<")
	var pause string
	fmt.Scanln(&pause)
}


func menu() {
	character := character.InitCharacter()

	for {
		fmt.Println("\n=== MENU PRINCIPAL ===")
		fmt.Println("1. Afficher les infos du personnage")
		fmt.Println("2. Aller chez le Forgeron")
		fmt.Println("3. Qui sont-ils")
		fmt.Println("4. Entrainement")

		fmt.Println("5. Quitter") 

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
			trainingFight(&character)
		case 5:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
