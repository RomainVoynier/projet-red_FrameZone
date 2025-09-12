package main

import (
	"fmt"
)

type Character struct {
	Nom               string
	Classe            string
	Niveau            int
	HpMax    int
	HpActuels int
	Inventaire        []string
}
