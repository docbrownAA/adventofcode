package game

import (
	"tuto/tpmoropionclasse/joueur"
	"tuto/tpmoropionclasse/plateau"
)

type Game struct {
	Plateau *plateau.Plateau
	Joueur1 *joueur.Joueur
	Joueur2 *joueur.Joueur
}

func Init(Plateau plateau.Plateau, Joueur1 *joueur.Joueur, Joueur2 *joueur.Joueur) Game {
	return Game{Plateau: &Plateau, Joueur1: Joueur1, Joueur2: Joueur2}
}
