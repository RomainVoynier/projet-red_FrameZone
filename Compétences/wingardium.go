package wingardium

import (
	"fmt"
)

type Personnage struct {
	Nom    string
	Skills []string
}

func spellBook(p *Personnage) {
	sort := "Boule de feu"

	for _, skill := range p.Skills {
		if skill == sort {
			fmt.Println("Le sort « Boule de feu » est déjà appris.")
			return
		}
	}

	p.Skills = append(p.Skills, sort)
	fmt.Println("Le sort « Boule de feu » a été ajouté à votre grimoire.")
}
