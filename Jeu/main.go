package main

import (
    "projet-red_FrameZone/Jeu/combat"
    "projet-red_FrameZone/Jeu/equipement"
    "projet-red_FrameZone/Jeu/menu"
)


func main() {
    fmt.Println("Bienvenue dans le jeu Projet Red FrameZone !")

    // Initialisation du personnage joueur
    player := character.InitPlayer() // Si InitPlayer est dans character

    // Lancer un combat d'entraînement
    combat.TrainingFight(player) // Appelle la fonction trainingFight avec player

    // Ici tu peux lancer d'autres choses : menu, exploration, etc.

    fmt.Println("Merci d'avoir joué ! À bientôt !")
}
