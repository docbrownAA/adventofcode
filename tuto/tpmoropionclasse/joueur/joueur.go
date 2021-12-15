package joueur

import "fmt"

type Joueur struct {
	Nom     string
	Symbole string
	AJoue   bool
}

func New(Nom string, Symbole string) Joueur {
	return Joueur{Nom, Symbole, false}
}

func (j *Joueur) Affichage() {
	fmt.Println("--- ", j.Nom, ": ", j.Symbole, "--- a joué:", j.AJoue)
}
