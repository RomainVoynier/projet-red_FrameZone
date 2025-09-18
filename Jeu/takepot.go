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
	potionIndex := -1
	for i, item := range j.Inventaire {
		if item == "potion" {
			potionIndex = i
			break
		}
	}

	if potionIndex == -1 {
		fmt.Println("Aucune potion disponible dans l'inventaire.")
		return
	}

	// Retirer la potion de l'inventaire
	j.Inventaire = append(j.Inventaire[:potionIndex], j.Inventaire[potionIndex+1:]...)

	// Soigner le joueur
	soin := 50
	ancienPV := j.PV
	j.PV += soin
	if j.PV > j.PVMax {
		j.PV = j.PVMax
	}

	fmt.Printf("Vous avez utilisé une potion ! (+%d PV)\n", j.PV-ancienPV)
	fmt.Printf("Points de vie : %d / %d\n", j.PV, j.PVMax)
}
