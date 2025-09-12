package main

import (
	"fmt"
)

type Character struct {
	Nom               string
	Class             string
	Level             int
	HpMax    		  int
	HpActual 		  int
	Inventory       []string
}
