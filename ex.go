package main

import (
	"fmt"
)

type Voiture struct {
	model string
	body  string
	power int
}

func main() {
	//Initialisation de mes voitures
	var v1 Voiture
	var v2 Voiture

	//Initialisation des paramètres de ma voiture 1
	v1.model = "Audi TT"
	v1.body = "Coupé"
	v1.power = 245

	//Initialisation des paramètres de ma voiture 2
	v2.model = "Mercedes-AMG"
	v2.body = "Coupé"
	v2.power = 585

	v1.affichage()
	v2.affichage()

	v1.gainPower()
	v1.gainPower()
	v1.gainPower()

	v1.affichage()

}

func (v Voiture) affichage() {
	fmt.Println("Model :", v.model)
	fmt.Println("Carrosserie :", v.body)
	fmt.Println("Puissance :", v.power)
	fmt.Println("-----------------------")
}

func (v *Voiture) gainPower() {
	v.power += 10
	fmt.Println(v.model, ": Nouvelle puissance :", v.power)
	fmt.Println("-----------------------")
}
