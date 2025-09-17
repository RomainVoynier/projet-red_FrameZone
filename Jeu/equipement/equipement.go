package equipement

import "fmt"

// Boutique d'équipements disponibles
var Boutique = []Objet{
	{"Couronne de Lauriers", 5, "Tete", 5},
	{"Tronc d'Arbre", 15, "Torse", 20},
	{"Bottes de Sapin", 10, "Pieds", 10},
}

// Structure Equipement
type Objet struct {
	Nom     string
	Cout    int
	Slot    string // "Tete", "Torse", "Pieds"
	BonusHP int    // Exemple de bonus (on peut en ajouter d'autres plus tard)
}
type Equipement struct {
    Tete  *ItemEquipement
    Torse *ItemEquipement
    Pieds *ItemEquipement
}

type ItemEquipement struct {
    Nom     string
    Slot    string
    Cout    int
    BonusHP int
}



// Menu du forgeron : Achat d'équipement
func ForgeronMenu(c *character.Character) {
	for {
		fmt.Println("\nBienvenue chez le Forgeron")
		fmt.Printf("Smic actuel : %d\n", c.Smic)
		fmt.Println("Objets disponibles :")
		for i, obj := range Boutique {
			fmt.Printf("%d. %s (%s) : %d Smic, +%d HP\n", i+1, obj.Nom, obj.Slot, obj.Cout, obj.BonusHP)
		}
		fmt.Println("0. Retour")

		var choix int
		fmt.Print("Choix : ")
		_, err := fmt.Scanln(&choix)
		if err != nil || choix < 0 || choix > len(Boutique) {
			fmt.Println("Choix invalide.")
			continue
		}

		if choix == 0 {
			fmt.Println("Retour au menu principal.")
			return
		}

		objet := Boutique[choix-1]
		if c.Smic < objet.Cout {
			fmt.Println("Trop pauvre pour cet objet.")
			continue
		}

		// Achat
		c.Smic -= objet.Cout

		// Équipement selon le slot
		switch objet.Slot {
		case "Tete":
			c.Equipement.Tete = &objet
		case "Torse":
			c.Equipement.Torse = &objet
		case "Pieds":
			c.Equipement.Pieds = &objet
		}

		// Mise à jour des HP max et actuels
		c.HpMax = c.CalculerHpMax()
		if c.HpActual > c.HpMax {
			c.HpActual = c.HpMax
		}

		fmt.Printf("Vous avez équipé %s ! Nouveau HP max : %d\n", objet.Nom, c.HpMax)
	}
}
