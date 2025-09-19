package main

import (
	"fmt"
)

type Joueur struct {
	Nom        string
	PV         int
	PVMax      int
	Inventaire []string
}

func takePot(j *Joueur) {
	// Vérifie si au moins une potion est présente
	potionTrouvee := false
	for _, item := range j.Inventaire {
		if item == "potion" {
			potionTrouvee = true
			break
		}
	}

	if !potionTrouvee {
		fmt.Println("Aucune potion disponible dans l'inventaire.")
		return
	}

	// Soigner le joueur sans retirer la potion
	soin := 50
	ancienPV := j.PV
	j.PV += soin
	if j.PV > j.PVMax {
		j.PV = j.PVMax
	}

	fmt.Printf("Vous avez utilisé une potion ! (+%d PV)\n", j.PV-ancienPV)
	fmt.Printf("Points de vie : %d / %d\n", j.PV, j.PVMax)
}
