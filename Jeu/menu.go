package main

import "fmt"

func main() {
	menu()
}

func afficherNoms() {
	fmt.Println("\n=== Liste des noms ===")
	fmt.Println("1. ABBA")
	fmt.Println("2. Steven Spielberg")
	fmt.Println("\n(Appuyez sur Entrée pour continuer)")
	var pause string
	fmt.Scanln(&pause)
}

func training() {
	fmt.Println("\n>>> Début de l'entraînement <<<")
	var pause string
	fmt.Scanln(&pause)
}

func menu() {
	character := InitCharacter() // directement ici, sans prefixe

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
			ForgeronMenu(&character) // assure-toi que ForgeronMenu est défini dans ce package
		case 3:
			afficherNoms()
		case 4:
			trainingFight(&character) // idem pour trainingFight
		case 5:
			fmt.Println("Au revoir.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
