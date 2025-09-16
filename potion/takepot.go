package potion

import (
	"fmt"
)

type Personnage struct {
	Nom        string
	PV         int
	PVMax      int
	Inventaire []string
}

func takePot(p *Personnage) {
	potionIndex := -1
	for i, item := range p.Inventaire {
		if item == "potion" {
			potionIndex = i
			break
		}
	}

	if potionIndex == -1 {
		fmt.Println(" Aucune potion disponible dans l'inventaire.")
		return
	}

	p.Inventaire = append(p.Inventaire[:potionIndex], p.Inventaire[potionIndex+1:]...)

	soin := 50
	ancienPV := p.PV
	p.PV += soin
	if p.PV > p.PVMax {
		p.PV = p.PVMax
	}

	fmt.Printf(" Vous avez utilisé une potion ! (+%d PV)\n", p.PV-ancienPV)
	fmt.Printf(" Points de vie : %d / %d\n", p.PV, p.PVMax)
}
