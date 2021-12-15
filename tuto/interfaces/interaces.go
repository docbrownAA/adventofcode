package interfaces

import (
	"fmt"
	"math"
)

type Forme interface { //Création de l'interface Forme
	Aire() float64      // signature de la méthode Aire
	Perimetre() float64 // signature de la méthode Perimetre
}

type Rectangle struct {
	Largeur  float64
	Longueur float64
}

type Cercle struct {
	rayon float64
}

func (r Rectangle) Aire() float64 {
	return r.Largeur * r.Longueur
}

func (r Rectangle) Perimetre() float64 {
	return (r.Largeur + r.Longueur) * 2
}

func (c Cercle) Aire() float64 {
	return math.Pi * c.rayon * c.rayon
}

func (c Cercle) Perimetre() float64 {
	return 2 * math.Pi * c.rayon
}

func AirePerimetrePresentation(f Forme) {
	fmt.Println("- Aire: ", f.Aire())
	fmt.Println("- Perimetre: ", f.Perimetre())
}

func Main() {
	var f Forme = Rectangle{5.0, 4.0} // Initialisation de l'interface forme // affectation de la structure Rectangle à l'interface Forme
	/*
		r := Rectangle{5.0, 4.0}
		fmt.Println("Type de f: ", f)
		fmt.Printf("Valeur de f: %v\n", f)
		fmt.Println("Aire du rectangle f: ", f.Aire())
		fmt.Println("f == r ? ", f == r) */
	var c Forme = Cercle{5.0}
	AirePerimetrePresentation(f)
	AirePerimetrePresentation(c)

}
