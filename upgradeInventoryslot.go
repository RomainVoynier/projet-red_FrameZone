package main

import (
	"fmt"
)
func accessInventory(inventory []string) {
    fmt.Println("\n Inventaire :")
    if len(inventory) == 0 {
        fmt.Println("Ton inventaire est vide.")
        return
    }
    for i, item := range inventory {
        fmt.Printf("%d. %s\n", i+1, item)
    }
}


